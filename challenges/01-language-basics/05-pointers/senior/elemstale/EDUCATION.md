# Element pointers invalidated by append

## Intuition

A reallocating append detaches pointers taken beforehand; use the current slice's index or re-take the address afterward.

## Approach

1. After the slice header is reassigned/reallocated, an old element pointer `p` is stale.
2. Re-take `p = &s[0]` against the current backing array before writing.

## Solution

```go
func FirstOf(xs []int) []int {
	s := make([]int, len(xs), len(xs)) // len == cap, so append reallocates
	copy(s, xs)
	p := &s[0]
	s = append(s, 0)
	p = &s[0]
	*p = 42
	return s
}
```

## Walkthrough

Writing through the stale `p` misses the live element; re-deriving `p` from `s[0]` targets the current storage so `*p = 42` lands.

## Pitfalls

- After a reallocating append, `&s[0]` from before is stale.
- Index the current slice (`s[0]`) or re-take the pointer.
