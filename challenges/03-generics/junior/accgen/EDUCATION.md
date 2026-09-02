# Accumulator

## Intuition

Keeping the sum in `T` preserves the caller's type and precision, while the mean is inherently fractional — so only that method converts.

## Approach

1. `Add`: add to `total`, increment `n`.
2. `Sum`: return `total`.
3. `Mean`: return 0 when `n` is zero, else the converted quotient.

## Solution

```go
func (a *Acc[T]) Add(v T) {
	a.total += v
	a.n++
}

func (a *Acc[T]) Sum() T {
	return a.total
}

func (a *Acc[T]) Mean() float64 {
	if a.n == 0 {
		return 0
	}
	return float64(a.total) / float64(a.n)
}
```

## Walkthrough

`Add(1); Add(2); Mean()` computes `3.0 / 2.0`, giving `1.5` rather than the truncated `1`.

## Pitfalls

- Computing `a.total / T(a.n)` and converting afterwards.
- Dividing by zero before any value is added.
- Storing the total as `float64` and losing the caller's integer type in `Sum`.
