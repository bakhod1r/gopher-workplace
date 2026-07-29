# Rounding modes and bias

## Intuition

`math.Round` rounds halves **away from zero** (2.5→3, 3.5→4), which biases sums
upward when many values end in .5. Banker's rounding sends ties to the **even**
neighbor (2.5→2, 3.5→4), so the bias cancels over a data set:

```go
math.RoundToEven(x)
```

## Approach

1. Bug: math.Round rounds half AWAY from zero (2.5->3, -2.5->-3).
2. Bankers rounding sends ties to the even neighbor.
3. Fix: use math.RoundToEven(x).

## Solution

```go
import "math"

func Round(x float64) float64 {
	return math.RoundToEven(x)
}
```

## Walkthrough

2.5 -> nearest evens 2 and 4, pick even 2. 3.5 -> pick even 4. 2.6 not a tie -> 3.

## Pitfalls

- `math.Round` ≠ `math.RoundToEven`; pick deliberately.
- Half-to-even is the IEEE-754 default rounding for float ops.
- Rounding to N decimals still needs scaling; the tie rule is the subtle part.
