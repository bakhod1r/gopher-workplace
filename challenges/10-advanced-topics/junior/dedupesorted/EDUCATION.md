# Collapse Runs Without New Memory

## Intuition

Sorting has already grouped the duplicates. One pass with a write cursor rewrites the survivors to the front, and the prefix is the result.

## Approach

1. Return early for an empty slice.
2. Start `k` at 1 — the first element always survives.
3. For each later element, keep it when it differs from `s[k-1]`.
4. Return `s[:k]`.

## Solution

```go
// Dedupe removes consecutive duplicates from the sorted slice s,
// reusing s's storage, and returns the deduplicated prefix.
//
// Elements past the returned length are unspecified.
//
// Examples:
//
// 	Dedupe([]int{1, 1, 2, 3, 3}) => []int{1, 2, 3}
func Dedupe(s []int) []int {
	if len(s) == 0 {
		return s
	}
	k := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[k-1] {
			s[k] = s[i]
			k++
		}
	}
	return s[:k]
}
```

## Walkthrough

[1 1 2 3 3 3]: i=1 equals s[0], skipped. i=2 (2) differs, written to s[1], k=2. i=3 (3) differs from s[1], written to s[2], k=3. The rest match s[2]. Result s[:3] = [1 2 3].

## Pitfalls

- Comparing `s[i]` with `s[i-1]` — correct only until the array has been overwritten behind you.
- Starting `k` at 0 and dropping the first element.
