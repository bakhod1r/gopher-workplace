# Circuit Breaker

## Intuition

A breaker converts a slow cascading failure into a fast local one. The load-shedding only works if the open state short-circuits *before* the dependency is touched.

## Approach

1. While open and inside the cooldown, return `ErrOpen` without calling `op`.
2. Otherwise run the operation.
3. On error, increment the failure count and open when it reaches `Threshold`, stamping `openedAt`.
4. On success, reset the count and close.

## Solution

```go
func (b *Breaker) Call(op Op) error {
	if b.open {
		if b.clock.Now().Sub(b.openedAt) < b.Cooldown {
			return ErrOpen
		}
	}

	err := op.Do()
	if err != nil {
		b.failures++
		if b.failures >= b.Threshold {
			b.open = true
			b.openedAt = b.clock.Now()
		}
		return err
	}

	b.failures = 0
	b.open = false
	return nil
}

func (b *Breaker) IsOpen() bool { return b.open }
```

## Walkthrough

A failed probe re-stamps `openedAt`, so the next call is `ErrOpen` again — the cooldown restarts rather than allowing a probe storm.

## Pitfalls

- Calling `op.Do()` first and checking the state afterwards, which defeats the whole point.
- Not resetting `failures` on success, so unrelated intermittent errors eventually trip the breaker.
- Leaving `openedAt` unchanged on a failed probe, letting every subsequent call probe again.
