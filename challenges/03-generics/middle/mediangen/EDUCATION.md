# Median

## Intuition

Sorting a clone costs an allocation but keeps the function pure, which matters when the same sample window feeds several statistics.

## Approach

1. Return zero and `false` for an empty slice.
2. Sort a clone.
3. Return the middle element, or the average of the two middle elements.

## Solution

```go
func Median[T Float](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	c := slices.Clone(s)
	slices.Sort(c)
	mid := len(c) / 2
	if len(c)%2 == 1 {
		return c[mid], true
	}
	return (c[mid-1] + c[mid]) / 2, true
}
```

## Walkthrough

`Median([]float64{1, 2, 3, 4})` averages `c[1]` and `c[2]`, giving `2.5`.

## Pitfalls

- Sorting the caller's slice.
- Using `len/2` for both middles in the even case.
- Rounding the even-case average, which the spec does not ask for.
