# Ordering That Ignores The Length Rule

## Intuition

`slices.Compare` is dictionary order: it walks both slices, returns on the first differing element, and only uses length when the shorter is a prefix of the longer. `[9]` versus `[1 2]` differs at index 0, so it answers +1 and length never enters into it.

## Approach

1. Compare the lengths first and return that result when they differ.
2. Delegate to `slices.Compare` only for equal lengths, where lexicographic order is exactly the tie-break wanted.

## Solution

```go
func ComparePaths(a, b []int) int {
	if len(a) != len(b) {
		return cmp.Compare(len(a), len(b))
	}
	return slices.Compare(a, b)
}
```

## Walkthrough

`ComparePaths([9], [1,2])` returns +1, so the two-segment path ranks ahead of the one-segment path — the opposite of the documented rule.

## Pitfalls

- Assuming `slices.Compare` is "shortest first" because it does compare lengths — it does so last, not first.
- Returning `len(a) - len(b)` as the result; the sign is right but the contract says -1, 0, or +1.
