# Nightly Batch Cut-off

## Intuition

`WithDeadline` is the primitive; `WithTimeout` is sugar over it. Scheduler-driven jobs think in absolute instants, so `WithDeadline` is the natural fit. Handing it an instant that has already gone by makes the context finish at construction time — deterministic, no sleeping, no clock assertions.

## Approach

1. `ctx, cancel := context.WithDeadline(context.Background(), cutoff)`.
2. `defer cancel()`.
3. `<-ctx.Done()`.
4. Return `ctx.Err()`.

## Solution

```go
import "context"

// MissedCutoff builds the batch job's context from the absolute cut-off instant
// handed down by the scheduler, waits for that context to finish, and returns
// the reason. Callers pass a cut-off that has already passed, meaning the job
// started too late and must refuse to run.
//
// Examples:
//
//	MissedCutoff(time.Now().Add(-time.Hour)) => context.DeadlineExceeded
//	MissedCutoff(time.Unix(0, 0))            => context.DeadlineExceeded
//	errors.Is(..., context.Canceled)         => false
func MissedCutoff(cutoff time.Time) error {
	ctx, cancel := context.WithDeadline(context.Background(), cutoff)
	defer cancel()
	<-ctx.Done()
	return ctx.Err()
}
```

## Walkthrough

- The deadline is compared against the current time when the context is built.
- Because `cutoff` is in the past, the context is created already expired and `Done()` is closed.
- `Err()` returns `context.DeadlineExceeded` for every past instant, including the zero `time.Time`.

## Pitfalls

- Deriving never *extends* a deadline: a child deadline later than the parent's is ignored.
- `time.Time{}` (the zero value) is year 1, firmly in the past; it is not treated as "no deadline".
- Still call `cancel()`. An expired context holds resources like any other.
