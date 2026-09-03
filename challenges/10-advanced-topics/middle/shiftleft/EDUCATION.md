# Move Elements Over Themselves

## Intuition

A left shift moves every element to a lower index, so a forward copy never overwrites something it has not yet read. `copy` is specified to handle that, which is why no temporary is needed.

## Approach

1. Clamp `n`.
2. `k := copy(s, s[n:])`.
3. Return `s[:k]`.

## Solution

```go
// Shift drops the first n elements of s by moving the rest to the front,
// in place, and returns the shortened slice.
//
// The source and destination ranges overlap, which copy handles correctly.
// n is clamped into [0, len(s)].
//
// Examples:
//
// 	Shift([]int{1, 2, 3, 4}, 2) => []int{3, 4}
func Shift(s []int, n int) []int {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	k := copy(s, s[n:])
	return s[:k]
}
```

## Walkthrough

For [1 2 3 4] and n = 2: `copy` moves 3 into index 0 and 4 into index 1, returning 2. The result is `s[:2]` = [3 4].

## Pitfalls

- Copying into a temporary first — correct and pointless for a left shift.
- Returning `s[:len(s)-n]` without clamping, which is negative for a large `n`.
