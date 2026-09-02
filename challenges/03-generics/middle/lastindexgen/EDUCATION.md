# Last Index

## Intuition

Scanning backwards avoids a "best so far" variable and stops as soon as the answer is known — for a trailing separator that is usually one iteration.

## Approach

1. Count `i` down from the last index.
2. Return `i` at the first match.
3. Return `-1` afterwards.

## Solution

```go
func LastIndex[T comparable](s []T, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}
```

## Walkthrough

`LastIndex([]int{7,1,7}, 7)` matches at index 2 immediately and never inspects the rest.

## Pitfalls

- Scanning forwards and overwriting a saved index — correct but always O(n).
- Starting at `len(s)`, which is out of range.
- Returning `0` for "not found".
