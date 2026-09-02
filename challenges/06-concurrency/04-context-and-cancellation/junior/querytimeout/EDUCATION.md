# Query Timeout Budget

## Intuition

A timeout is just a deadline computed as `now + d`. With `d == 0` the deadline is *now*, so the context is finished the moment it is built and `Err()` is `DeadlineExceeded`. Nothing waits on a real clock, which is what makes this test deterministic under `-race` and on a slow CI box.

## Approach

1. `ctx, cancel := context.WithTimeout(context.Background(), 0)`.
2. `defer cancel()` — required even for an already-expired context, to release its timer.
3. `<-ctx.Done()` so the state is settled.
4. Return `ctx.Err()`.

## Solution

```go
// ExhaustedQueryBudget models a database query whose time budget is already
// spent by the time the query is dispatched: the surrounding request used it
// all up. It builds the query context with a zero timeout, waits for it to
// finish, and returns the reason.
//
// The timeout must be zero or negative so the context is done immediately —
// never depend on wall-clock time passing.
//
// Examples:
//
//	ExhaustedQueryBudget()                                     => context.DeadlineExceeded
//	errors.Is(ExhaustedQueryBudget(), context.DeadlineExceeded) => true
//	errors.Is(ExhaustedQueryBudget(), context.Canceled)         => false
func ExhaustedQueryBudget() error {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()
	<-ctx.Done()
	return ctx.Err()
}
```

## Walkthrough

- `WithTimeout(parent, 0)` computes a deadline of `time.Now()`, which is not in the future.
- The context is therefore created already expired and `Done()` is closed.
- `<-ctx.Done()` returns immediately and `Err()` reports `context.DeadlineExceeded` — which the query layer maps to a 504 rather than a 499.

## Pitfalls

- Asserting on elapsed wall-clock time instead of the error value is what made the original test flaky.
- `context.Canceled` and `context.DeadlineExceeded` are different values, and the deadline latches its error first — a later `cancel()` cannot overwrite it.
- Skipping `defer cancel()` still leaks the timer even when the deadline has passed; `go vet` flags it.
