# Rounding modes and bias

## The idea

`math.Round` rounds halves **away from zero** (2.5→3, 3.5→4), which biases sums
upward when many values end in .5. Banker's rounding sends ties to the **even**
neighbor (2.5→2, 3.5→4), so the bias cancels over a data set:

```go
math.RoundToEven(x)
```

## Why it matters

Finance, statistics, and IEEE-754 default arithmetic all use round-half-to-even
precisely to avoid cumulative drift. Choosing the wrong mode is a real accounting
discrepancy that grows with volume.

## Watch out

- `math.Round` ≠ `math.RoundToEven`; pick deliberately.
- Half-to-even is the IEEE-754 default rounding for float ops.
- Rounding to N decimals still needs scaling; the tie rule is the subtle part.
