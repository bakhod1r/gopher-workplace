# Conditional counting

## Intuition

A running counter incremented under an `if` inside a loop is the basis of filters, histograms, and predicates.

## Approach

1. Range the slice.
2. Count elements equal to `target`.

## Solution

```go
func CountEqual(xs []int, target int) int {
	count := 0
	for _, v := range xs {
		if v == target {
			count++
		}
	}
	return count
}
```

## Walkthrough

`[1 2 2 3 2]` has three 2s → 3.

## Pitfalls

- Return 0 for no matches, never -1 or a sentinel.
- A nil slice ranges zero times.
