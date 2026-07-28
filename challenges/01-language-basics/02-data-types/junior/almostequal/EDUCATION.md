# Comparing floats with a tolerance

## The idea

Binary floats can't represent most decimals exactly, so `==` is the wrong tool.
Compare within a small tolerance instead:

```go
math.Abs(a-b) <= eps
```

## Why it matters

`0.1 + 0.2 != 0.3` in float64. Any equality check on *computed* floats needs a
tolerance, or it fails unpredictably.

## Watch out

- If `a` or `b` is NaN, `a-b` is NaN and the comparison is false — NaN is never
  "almost equal".
- Absolute tolerance suits values near the same magnitude; relative tolerance
  scales across magnitudes.
- The right `eps` depends on the operations performed, not a universal constant.
