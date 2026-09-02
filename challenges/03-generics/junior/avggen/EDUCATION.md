# Average

## Intuition

Summing in `T` and converting at the end truncates for integer instantiations and can overflow. Converting each element first keeps one implementation correct for all members of the set.

## Approach

1. Return `0` for an empty slice.
2. Accumulate `float64(v)` into a `float64` total.
3. Divide by `float64(len(s))`.

## Solution

```go
func Average[T Number](s []T) float64 {
	if len(s) == 0 {
		return 0
	}
	var total float64
	for _, v := range s {
		total += float64(v)
	}
	return total / float64(len(s))
}
```

## Walkthrough

`Average([]int{1, 2, 3})` converts each element, totals `6.0`, and divides by `3.0` to give `2`.

## Pitfalls

- Dividing by `len(s)` without converting, which does not compile.
- Summing into a `T` and converting after — `[]int{1, 2}` would then average to `1`, not `1.5`.
- Dividing by zero on an empty slice.
