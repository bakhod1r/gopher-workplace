# Circuit Breaker

## Intuition

A breaker protects a failing dependency from being hammered. It counts
*consecutive* failures — a single success means the trouble has passed and the
count resets. Once the count reaches the threshold the breaker opens, and from
then on calls fail fast without ever reaching the wrapped function.

## Approach

1. If the breaker is open, return an error immediately — do not call `fn`.
2. Otherwise call `fn`.
3. On error: increment the counter and open if it reached the threshold.
4. On success: reset the counter to 0.

## Solution

```go
func (b *Breaker) Call(fn func() error) error {
	if b.IsOpen {
		return errors.New("circuit open")
	}
	if err := fn(); err != nil {
		b.ConsecutiveFails++
		if b.ConsecutiveFails >= b.Threshold {
			b.IsOpen = true
		}
		return err
	}
	b.ConsecutiveFails = 0
	return nil
}
```

## Walkthrough

With `Threshold: 2`:

| call | before | `fn` run? | after |
|------|--------|-----------|-------|
| `failFn` | closed, fails 0 | yes | fails 1, still closed |
| `failFn` | closed, fails 1 | yes | fails 2, **open** |
| `okFn` | open | no | unchanged, returns `"circuit open"` |

The third assertion is the important one: the breaker returns the sentinel error
even though the supplied function would have succeeded. That is fail-fast.

## Pitfalls

- **Checking `IsOpen` after calling `fn`.** The dependency still gets hit — the
  breaker becomes decoration.
- **`>` instead of `>=`.** With `Threshold: 2` the breaker would need three
  failures to open.
- **Not resetting on success.** The counter then measures total failures, not
  consecutive ones, and the breaker eventually trips on a healthy service.
- **Returning a fresh error object where the caller compares strings.** The test
  compares `err.Error()`, so the message text is part of the contract.

## What a real breaker adds

Production breakers have a half-open state: after a cooldown they let one probe
request through and close again if it succeeds. This puzzle stops at the
open/closed core so the state transitions stay visible.
