# Index Of

## Intuition

Returning early is what makes this the first index. If you kept scanning and overwrote a saved index, you would end up with the last match instead.

## Approach

1. Range with both index and value.
2. Return `i` at the first match.
3. Return `-1` after the loop.

## Solution

```go
func IndexOf[T comparable](s []T, v T) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}
```

## Walkthrough

`IndexOf([]int{5, 7, 7}, 7)` matches at index 1 and returns straight away, so the second `7` at index 2 never wins.

## Pitfalls

- Returning `0` for "not found" — `0` is a valid index.
- Recording the index in a variable and returning it after the loop, which yields the last match.
- Using `[T any]`, which cannot use `==`.
