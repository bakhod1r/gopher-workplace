# Growing a caller's slice

## Intuition

Because append returns a possibly-new header, mutating the caller's slice requires assigning the result through a `*[]int`.

## Approach

1. `sp` is a `*[]int` — a pointer to the caller's slice header.
2. `*sp = append(*sp, v)` writes the (possibly relocated) new header back.
3. Without the pointer, `append`'s new header would be lost.

## Solution

```go
func Push(sp *[]int, v int) {
	*sp = append(*sp, v)
}
```

## Walkthrough

`Push(&xs, 1)` on an empty slice: `append` returns a length-1 slice; storing it through `*sp` updates the caller's `xs`.

## Pitfalls

- A plain `[]int` parameter can't propagate a reallocated header.
- `*sp = append(*sp, v)` updates the caller.
