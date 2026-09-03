# Concatenate Slices With One Allocation

## Intuition

`append` grows by guessing because it cannot see what is coming. When the caller can, one counting pass removes every reallocation and every intermediate copy.

## Approach

1. Sum the parts' lengths.
2. Return nil when the total is zero.
3. `make([]int, 0, n)` and append the parts into it.

## Solution

```go
// AppendAll returns every part concatenated in order.
//
// The final length is known before the first append, so the result should
// be allocated once instead of growing through every doubling.
//
// Examples:
//
// 	AppendAll([][]int{{1}, {2, 3}}) => []int{1, 2, 3}
func AppendAll(parts [][]int) []int {
	n := 0
	for _, p := range parts {
		n += len(p)
	}
	if n == 0 {
		return nil
	}
	out := make([]int, 0, n)
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}
```

## Walkthrough

64 parts of 4 elements is 256. Starting from nil, `append` reallocates at 1, 2, 4 … 256 and copies 255 elements along the way; the sized version allocates once and copies each element exactly once.

## Pitfalls

- `make([]int, n)` instead of `make([]int, 0, n)`, which prepends n zeros.
- Returning `parts[0]` when there is only one part — that shares storage with the caller.
