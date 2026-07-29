# Modular wraparound

## Intuition

Clock arithmetic is arithmetic modulo 24. Adding hours and taking `% 24` folds
any value back into range — almost. Because Go's `%` follows the dividend's
sign, a negative sum yields a negative remainder, so you normalize:

```go
((h+add)%24 + 24) % 24 // always 0..23
```

The `+24` pushes a negative remainder positive; the final `%24` trims a value
that was already in range.

## Approach

1. Compute h+add.
2. Take % 24 (may be negative in Go).
3. Add 24 and take % 24 again to force a non-negative 0..23 result.

## Solution

```go
func AddHours(h, add int) int {
	return ((h+add)%24 + 24) % 24
}
```

## Walkthrough

AddHours(0,-1): (0-1)%24 = -1, (-1+24)%24 = 23.

## Pitfalls

- `-1 % 24 == -1`, not 23 — the source of the bug.
- The double-mod is needed only when the input can be negative.
- For powers of two you can mask (`x & 23` only if 24 were a power of two — it
  is not, so use `%`).
