# Batch A Slice Into Fixed-Size Windows

## Intuition

Splitting is a question of where the boundaries are, not of moving data. Every group can point into the original array; only the little slice of headers is new.

## Approach

1. Reject `n <= 0`.
2. Preallocate the outer slice with the ceiling-divided group count.
3. Step `i` by `n`, clamp `end` to `len(s)`, append `s[i:end]`.

## Solution

```go
// Chunk splits s into consecutive groups of at most n elements.
//
// The last group holds the remainder. For n <= 0 the result is nil. The
// groups are views into s — no element is copied.
//
// Examples:
//
// 	Chunk([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}
func Chunk(s []int, n int) [][]int {
	if n <= 0 {
		return nil
	}
	out := make([][]int, 0, (len(s)+n-1)/n)
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		out = append(out, s[i:end])
	}
	return out
}
```

## Walkthrough

For five elements and n = 2, the group count is (5+1)/2 = 3. The windows are [0:2], [2:4] and [4:5] — the last one clamped.

## Pitfalls

- `s[i : i+n]` without the clamp panics on the last group.
- Copying each group into a fresh slice — correct output, wrong memory behaviour.
