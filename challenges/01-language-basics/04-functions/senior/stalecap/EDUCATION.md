# Element pointers and reallocation

## Intuition

`append` past capacity moves the backing array; any pointer/slice into the old array is now detached from the live slice.

## Approach

1. After the slice grows (reallocates), an old element pointer is stale.
2. Re-take `p = &xs[0]` against the new backing array before writing.

## Solution

```go
func FirstAfterGrow(v int) int {
	xs := make([]int, 1, 1) // len==cap==1, so append reallocates
	xs[0] = 10
	p := &xs[0]
	xs = append(xs, v)
	p = &xs[0]
	*p = 99
	return xs[0]
}
```

## Walkthrough

`append` may move the backing array, leaving `p` pointing at freed storage. Re-deriving `p` from `xs[0]` targets the live element so `*p = 99` sticks.

## Pitfalls

- After a reallocating append, old `&xs[i]` is stale.
- Index the current slice (`xs[0]`) or re-take the address.
