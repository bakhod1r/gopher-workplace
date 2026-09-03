# Binary Search Against The Grain

## Intuition

`BinarySearchFunc` does not check anything: it halves the range using the sign your comparator returns. Its precondition is that `cmp(s[i], target)` is non-decreasing in `i`. On a descending board `cmp.Compare(e.Score, target)` *decreases* with `i`, so the halving walks away from the answer.

## Approach

1. Compare `target` against `e.Score`, not the other way round, so the sign rises with the index.
2. Treat `ok == false` as absent and report -1.

## Solution

```go
func FindScore(board []Entry, score int) (int, bool) {
	i, ok := slices.BinarySearchFunc(board, score, func(e Entry, target int) int {
		return cmp.Compare(target, e.Score)
	})
	if !ok {
		return -1, false
	}
	return i, true
}
```

## Walkthrough

On `[9 5 1]` searching for 5, the buggy comparator at the midpoint `5` returns 0 and happens to work; searching for 9 it inspects `5`, gets `-1`, and moves right — away from the 9 sitting at index 0.

## Pitfalls

- Assuming a wrong comparator produces a panic or a wrong-but-close answer; it produces confident nonsense.
- Reversing the slice at every call to "fix" it, turning an O(log n) lookup into O(n).
