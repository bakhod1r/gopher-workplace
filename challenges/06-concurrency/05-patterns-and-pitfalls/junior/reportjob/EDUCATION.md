# Respecting a Context

## Intuition

A context is a cancellation signal with an explanation attached. Checking it
before starting work is the cheapest possible win: no goroutines, no locking,
no wasted CPU when the caller has already gone.

## Approach

1. `if err := ctx.Err(); err != nil { return 0, err }`.
2. Start one goroutine per row, each locking the mutex to add its value.
3. `wg.Wait()` and return the total with a nil error.

## Solution

```go
import (
	"context"
	"sync"
)

func RunReport(ctx context.Context, rows []int) (int, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		total int
	)

	for _, row := range rows {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()

			mu.Lock()
			total += row
			mu.Unlock()
		}(row)
	}

	wg.Wait()
	return total, nil
}
```

## Walkthrough

With a live context and rows 1, 2, 3, three goroutines each take the mutex
once and the total reaches 6. With a cancelled context, `ctx.Err()` returns
`context.Canceled` on the first line and the function returns `0, err` without
spawning anything.

## Pitfalls

- Returning a bespoke error instead of `ctx.Err()`, so callers cannot use `errors.Is(err, context.Canceled)`.
- Returning a partial total alongside the error — a cancelled report must not look like a complete one.
- Adding to `total` without the mutex, which the race detector flags immediately.
