# Integer vs float average

## The idea

The mean of integers is a *ratio*, so the division must happen in floating point.
Summing in `int` and dividing by the count as ints truncates the fraction:

```go
float64(sum) / float64(n) // 7/2 -> 3.5, not 3
```

## Why it matters

Averages, rates, and ratios are everywhere. Dividing before converting to float
(or dividing two ints) silently floors the result — a classic reporting bug.

## Watch out

- Convert to `float64` before dividing, not after.
- Guard `n == 0` to avoid a divide-by-zero panic.
- Summing many ints can overflow; widen the accumulator if needed.
