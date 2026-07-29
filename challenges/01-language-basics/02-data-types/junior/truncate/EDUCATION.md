# Float-to-int truncation

## Intuition

Converting a `float64` to an `int` **truncates toward zero** — it drops the
fraction, it does not round:

```go
int(3.9)  // 3
int(-3.9) // -3
```

## Approach

1. Convert the float64 to int with int(amount).
2. Go truncates toward zero, which is exactly the required whole-dollar behavior.

## Solution

```go
func WholePart(amount float64) int {
	return int(amount)
}
```

## Walkthrough

WholePart(-9.99): int(-9.99) drops the fraction toward zero -> -9.

## Pitfalls

- `int(x)` truncates toward zero, so negatives round *up* toward zero.
- Converting a float larger than the int range is implementation-defined —
  guard the domain.
- For "floor" use `math.Floor`; for nearest use `math.Round`.
