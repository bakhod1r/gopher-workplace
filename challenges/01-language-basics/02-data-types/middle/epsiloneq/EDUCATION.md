# Comparing floats

## The idea

Binary floats cannot represent most decimals exactly, so `0.1+0.2` is
`0.30000000000000004`. Compare with a tolerance:

```go
math.Abs(a-b) <= eps
```

## Why it matters

`==` on computed floats is almost always a bug. A tolerance-based compare is the
standard fix; the right `eps` depends on the magnitudes and the number of
operations.

## Watch out

- If `a` or `b` is NaN, `a-b` is NaN and `NaN <= eps` is false — so NaN is
  correctly reported unequal, for free.
- Absolute tolerance fails across very different magnitudes; relative tolerance
  (`|a-b| <= eps*max(|a|,|b|)`) scales better.
- `eps == 0` reduces to exact equality (fine for identical values).
