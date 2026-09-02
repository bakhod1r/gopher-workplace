# Sum If

## Intuition

Fusing the two operations avoids one allocation and one extra pass — the same result as `Sum(Filter(s, keep))` with less work.

## Approach

1. Declare `var total T`.
2. Add each element `keep` accepts.
3. Return the total.

## Solution

```go
func SumIf[T Number](s []T, keep func(T) bool) T {
	var total T
	for _, v := range s {
		if keep(v) {
			total += v
		}
	}
	return total
}
```

## Walkthrough

`SumIf([]int{1, 2, 3}, isEven)` adds only `2`, returning `2`.

## Pitfalls

- Adding every element and multiplying by a ratio afterwards.
- Allocating a filtered slice first, which the puzzle explicitly avoids.
- Starting the total at a literal `0` instead of `var total T`.
