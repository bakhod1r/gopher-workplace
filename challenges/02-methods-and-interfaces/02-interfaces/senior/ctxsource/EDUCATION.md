# Context-Aware Source

## Intuition

Cancellation is only useful if someone checks it. In a tight loop the check has to be non-blocking, which is what `select` with a `default` gives you.

## Approach

1. `RangeSource.Next` yields `1..N` then reports drained.
2. In `SumWithContext`, run a non-blocking `select` on `ctx.Done()` at the top of each iteration.
3. Return `0, ctx.Err()` on cancellation — the partial sum is not a valid result.
4. Otherwise read the next value and accumulate, returning `sum, nil` when drained.

## Solution

```go
func (r *RangeSource) Next() (int, bool) {
	if r.pos >= r.N {
		return 0, false
	}
	r.pos++
	return r.pos, true
}

func SumWithContext(ctx context.Context, src Source) (int, error) {
	sum := 0
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}

		v, ok := src.Next()
		if !ok {
			return sum, nil
		}
		sum += v
	}
}
```

## Walkthrough

`TestCancelMidScan` cancels after five reads of a 10M source. The check at the top of the loop catches it on the next iteration, so total reads stay tiny — that is what "promptly" means here.

## Pitfalls

- Checking cancellation only once before the loop — a mid-scan cancel is then ignored.
- A blocking `case <-ctx.Done():` with no `default`, which stalls the loop on a live context.
- Returning the partial sum with the error, which invites callers to use half a result.
