# Exponential Backoff

## Intuition

A backoff is a tiny state machine with one field. Every call does two things
that must not be confused: it *answers* with the delay for this attempt, and it
*advances* internal state for the next one. Mixing the order changes the whole
sequence.

## Approach

1. Snapshot `b.current` into a local — that is the answer.
2. Double `b.current`.
3. If the doubled value passed `b.max`, pin it to `b.max`.
4. Return the snapshot.

## Solution

```go
func (b *Backoff) Next() time.Duration {
	d := b.current
	b.current *= 2
	if b.current > b.max {
		b.current = b.max
	}
	return d
}
```

## Walkthrough

With `max = 5s`, `current` starts at `1s`.

| Call | returned | current after |
|------|----------|---------------|
| 1 | 1s | 2s |
| 2 | 2s | 4s |
| 3 | 4s | 5s (8s clamped) |
| 4 | 5s | 5s |
| 5 | 5s | 5s |

The clamp is applied to the *stored* value, not to the returned one, which is
why the sequence saturates instead of oscillating.

## Pitfalls

- **Value receiver.** `func (b Backoff) Next()` compiles and always returns
  `1s`: the doubling lands on a copy that is discarded at return.
- **Advance before snapshot.** Returning `b.current` after `b.current *= 2`
  skips the first delay entirely — the sequence starts at 2s.
- **Clamping the return instead of the field.** `current` then grows without
  bound and eventually overflows int64.

## `time.Duration` is not a struct

`time.Duration` is a named `int64`. Comparison, multiplication and clamping are
plain integer operations; the pretty `5s` you see in test output comes from its
`String` method.
