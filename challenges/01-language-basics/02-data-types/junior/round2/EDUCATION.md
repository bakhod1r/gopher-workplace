# Rounding floats to decimals

## The idea

`math.Round` only rounds to whole numbers. To keep N decimals, shift the decimal
point by multiplying by 10^N, round to an integer, then shift back:

```go
p := math.Pow(10, float64(places))
math.Round(x*p) / p
```

## Why it matters

Displaying prices, measurements, and percentages needs fixed decimals. Doing it
by scaling with `math.Round` is the standard approach — and it exposes that
binary floats cannot represent every decimal exactly.

## Watch out

- `math.Round` rounds **half away from zero** (`2.5 -> 3`, `-2.5 -> -3`), unlike
  banker's rounding.
- Because `1.005` is stored slightly below 1.005, `Round(1.005, 2)` can give
  1.00 — a float fact, not a bug.
- For exact decimal money, use integer minor units, not rounded floats.

## Try it yourself

```go
math.Round(2.5)   // 3
math.Round(-2.5)  // -3
math.Round(1.2345*100) / 100 // 1.23
```
