# Insertion Points That Drift

## Intuition

Each `slices.Insert` shifts everything after the insertion point one slot right. The second position was measured against the *pre-insert* slice, so it must be corrected by the number of marks already placed.

## Approach

1. Clone the input.
2. Track how many marks have been inserted so far.
3. Insert at `p + done`, then bump the counter.

## Solution

```go
func InsertMarks[T any](s []T, at []int, mark T) []T {
	out := slices.Clone(s)
	done := 0
	for _, p := range at {
		if p < 0 || p > len(s) {
			continue
		}
		out = slices.Insert(out, p+done, mark)
		done++
	}
	return out
}
```

## Walkthrough

For `[1 2 3 4]` and `at = [1 3]`: the first insert gives `[1 0 2 3 4]`; index 3 in that slice is now the `3`, so the second mark lands before it instead of before the `4`.

## Pitfalls

- Walking `at` in descending order to dodge the drift — it works, but only if the caller sorts, which the doc does not promise.
- Validating `p` against the growing `out` instead of the original `s`.
