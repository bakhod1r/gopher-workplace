# Float-to-int truncation

## The idea

Converting a `float64` to an `int` **truncates toward zero** — it drops the
fraction, it does not round:

```go
int(3.9)  // 3
int(-3.9) // -3
```

## Why it matters

Truncation vs rounding is a frequent source of off-by-one bugs in money,
indexing, and binning. Know which one the conversion gives you (truncation) and
round explicitly with `math.Round`/`math.Floor` when you need it.

## Watch out

- `int(x)` truncates toward zero, so negatives round *up* toward zero.
- Converting a float larger than the int range is implementation-defined —
  guard the domain.
- For "floor" use `math.Floor`; for nearest use `math.Round`.
