# The for-range loop

## Intuition

`range` over a slice returns `(index, value)` copies; discarding the index with `_` gives a clean value iteration.

## Approach

1. Range and accumulate each element.

## Solution

```go
func SumRange(xs []int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}
```

## Walkthrough

`SumRange([1 2 3 4])` adds to 10.

## Pitfalls

- `v` is a copy; writing to it does not change the slice.
- A nil slice ranges zero times — no special case needed.
