# Modular wraparound

## The idea

Clock arithmetic is arithmetic modulo 24. Adding hours and taking `% 24` folds
any value back into range — almost. Because Go's `%` follows the dividend's
sign, a negative sum yields a negative remainder, so you normalize:

```go
((h+add)%24 + 24) % 24 // always 0..23
```

The `+24` pushes a negative remainder positive; the final `%24` trims a value
that was already in range.

## Why it matters

Ring buffers, angles (mod 360), hash bucket indices, and calendars all wrap.
The normalize idiom is the portable way to get a non-negative modulus.

## Watch out

- `-1 % 24 == -1`, not 23 — the source of the bug.
- The double-mod is needed only when the input can be negative.
- For powers of two you can mask (`x & 23` only if 24 were a power of two — it
  is not, so use `%`).

## Try it yourself

```go
(-1%24 + 24) % 24 // 23
(48%24 + 24) % 24 // 0
```
