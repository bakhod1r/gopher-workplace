# Range value copies

## Intuition

The range loop's value variable is a per-iteration copy; mutating it never reaches the underlying element — only indexing does.

## Approach

1. `for _, v := range xs` gives a **copy**; mutating `v` doesn't touch the slice.
2. Write via index: `xs[i] = v * 2`.

## Solution

```go
func DoubleAll(xs []int) {
	for i, v := range xs {
		_ = i
		xs[i] = v * 2
	}
}
```

## Walkthrough

The bug doubles a throwaway `v`, leaving the slice unchanged. Assigning `xs[i]` mutates the backing array in place.

## Pitfalls

- `for _, v := range xs { v = ... }` changes nothing.
- Use `xs[i]` (or a pointer element) to mutate in place.
