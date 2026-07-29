# Rounding floats to decimals

## Intuition

`math.Round` only rounds to whole numbers. To keep N decimals, shift the decimal
point by multiplying by 10^N, round to an integer, then shift back:

```go
p := math.Pow(10, float64(places))
math.Round(x*p) / p
```

## Approach

1. Compute p = 10^places via math.Pow.
2. Multiply x by p, round to an integer with math.Round, divide back by p.

## Solution

```go
import "math"

func Round(x float64, places int) float64 {
	p := math.Pow(10, float64(places))
	return math.Round(x*p) / p
}
```

## Walkthrough

Round(3.14159, 2): p=100, 314.159 -> math.Round=314, 314/100=3.14.

## Pitfalls

- `math.Round` rounds **half away from zero** (`2.5 -> 3`, `-2.5 -> -3`), unlike
  banker's rounding.
- Because `1.005` is stored slightly below 1.005, `Round(1.005, 2)` can give
  1.00 — a float fact, not a bug.
- For exact decimal money, use integer minor units, not rounded floats.
