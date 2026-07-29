# Early return for all/any checks

## Intuition

"All" returns false at the first counterexample then true after the loop; "any" is the mirror. Returning true mid-loop answers the wrong question.

## Approach

1. "All" fails fast on the first counterexample.
2. The bug returns true on the first positive; instead return false on the first non-positive, and true after the loop.

## Solution

```go
func AllPositive(xs []int) bool {
	for _, v := range xs {
		if v <= 0 {
			return false
		}
	}
	return true
}
```

## Walkthrough

Returning true early accepts a slice after seeing one positive, ignoring later negatives. Failing on `v <= 0` and returning true only at the end is correct; an empty slice is vacuously true.

## Pitfalls

- "All P" ⇒ fail fast on `!P`, succeed after the loop.
- "Any P" ⇒ succeed fast on `P`, fail after the loop.
