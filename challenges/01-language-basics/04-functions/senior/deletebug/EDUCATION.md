# The ordered-delete idiom

## Intuition

Removing index i keeps the prefix `xs[:i]` and the suffix `xs[i+1:]`; using `xs[i:]` for the suffix keeps the element you meant to drop.

## Approach

1. To delete index `i`, append the tail starting **after** it.
2. The bug uses `xs[i:]` (keeps the element); use `xs[i+1:]`.

## Solution

```go
func RemoveAt(xs []int, i int) []int {
	return append(xs[:i], xs[i+1:]...)
}
```

## Walkthrough

`append(xs[:i], xs[i:]...)` re-includes the element at `i`. Using `xs[i+1:]` skips it, removing index `i`.

## Pitfalls

- Suffix after deletion starts at `i+1`, not `i`.
- This idiom mutates the backing array; copy first if the caller keeps xs.
