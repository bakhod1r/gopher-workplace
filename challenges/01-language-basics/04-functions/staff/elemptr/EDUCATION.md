# Element pointers across reallocation

## Intuition

`append` beyond capacity detaches pre-existing element pointers from the live slice; you must re-take the address after growth.

## Approach

1. After the slice grows, the earlier element pointer is stale.
2. Re-take `p = &xs[0]` on the current backing array before writing 42.

## Solution

```go
func BumpHead(v int) []int {
	xs := make([]int, 1, 1)
	xs[0] = 1
	p := &xs[0]
	xs = append(xs, v)
	p = &xs[0]
	*p = 42
	return xs
}
```

## Walkthrough

`append` may relocate the array, so the old `p` points at freed memory. Re-deriving `p` from `xs[0]` writes 42 into the live head.

## Pitfalls

- After a reallocating append, `&xs[i]` from before is stale.
- Re-take it, or index the current slice.
