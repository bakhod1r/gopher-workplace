# First Replica to Answer Wins

## Intuition

Racing replicas is only a win if the losers actually stop. A shared `context.WithCancel` turns "I got my answer" into a broadcast: one `cancel()` closes one channel and every in-flight replica query unwinds. The other half of the trick is the *buffered* result channel — an unbuffered one would leave the losers blocked on a send that nobody will ever receive, which is a goroutine leak dressed up as a fast read.

## Approach

1. Reject the empty replica set and an already-finished context up front.
2. Derive `fanCtx, cancel := context.WithCancel(ctx)`; `defer cancel()`.
3. Start one goroutine per replica, each sending `{index, row, err}` into a channel buffered to `len(replicas)`.
4. Receive up to `len(replicas)` answers; return the first one whose error is nil.
5. Otherwise store each error at its own index and return `failures[0]`.

## Solution

```go
import (
	"context"
	"errors"
)
// Replica reads the requested row from one database replica.
type Replica func(ctx context.Context) (string, error)

// ErrNoReplicas reports that the read was issued against an empty replica set.
var ErrNoReplicas = errors.New("no replicas configured")

// ReadFromReplicas asks every replica for the same row at once and returns the
// first answer that comes back. The replicas share one derived context, so as
// soon as a winner is found the losers are cancelled and stop burning replica
// CPU on a row nobody will read.
//
// If the request context is already finished it returns ctx.Err() without
// touching a replica. If every replica fails it returns the failure reported by
// replicas[0]; an empty replica set returns ErrNoReplicas.
//
// Examples:
//
//	ReadFromReplicas(ctx, nil)                       => "", ErrNoReplicas
//	ReadFromReplicas(ctx, [slow, fast])              => the fast replica's row
//	ReadFromReplicas(cancelled ctx, [fast])          => "", context.Canceled
func ReadFromReplicas(ctx context.Context, replicas []Replica) (string, error) {
	if len(replicas) == 0 {
		return "", ErrNoReplicas
	}
	if err := ctx.Err(); err != nil {
		return "", err
	}

	fanCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	type answer struct {
		index int
		row   string
		err   error
	}
	answers := make(chan answer, len(replicas))
	for i, replica := range replicas {
		go func(i int, replica Replica) {
			row, err := replica(fanCtx)
			answers <- answer{index: i, row: row, err: err}
		}(i, replica)
	}

	failures := make([]error, len(replicas))
	for range replicas {
		a := <-answers
		if a.err == nil {
			return a.row, nil
		}
		failures[a.index] = a.err
	}
	return "", failures[0]
}
```

## Walkthrough

- With one fast and two parked replicas, the fast answer arrives first, the function returns, and the deferred `cancel()` releases both parked goroutines — each observes `context.Canceled`.
- A replica that fails does not end the read: its error is recorded and the loop keeps receiving.
- When all replicas fail the loop drains every answer, and `failures[0]` makes the returned error a property of the replica list, not of the scheduler.
- An already-cancelled request returns `context.Canceled` before a single goroutine is started.

## Pitfalls

- An unbuffered `answers` channel leaks every losing goroutine forever.
- Returning without `cancel()` — or forgetting the `defer` on the early-return paths — leaves the losers running.
- `return "", err` from the first failure turns one flaky replica into a failed read; the point of a fan-out is that failures are survivable.
- Returning "the first error received" is scheduler-dependent and makes the test flaky.
