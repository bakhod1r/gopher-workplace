# Pointers into slice elements

## Intuition

`&xs[i]` addresses the actual backing-array slot, so a `[]*int` of these aliases the slice; writes through them mutate xs.

## Approach

1. Range by **index**, not value, so each `&xs[i]` addresses the element.
2. Taking `&v` of the loop copy would alias the copy, not the slice.
3. Append each element address.

## Solution

```go
func Pointers(xs []int) []*int {
	var ps []*int
	for i := range xs {
		ps = append(ps, &xs[i])
	}
	return ps
}
```

## Walkthrough

`Pointers(xs)` collects `&xs[0], &xs[1], &xs[2]`; writing through `ps[1]` mutates `xs[1]` directly.

## Pitfalls

- Use `&xs[i]` (indexed), not `&v` from a range value in a way that detaches — though Go 1.22 makes `&v` per-iteration too.
- Appending to xs later may reallocate and detach these pointers.
