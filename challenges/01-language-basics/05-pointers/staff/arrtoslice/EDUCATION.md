# Slicing array pointers

## Intuition

`(&arr)[:]` produces a slice over the entire array, aliasing it; a shorter slice expression drops elements.

## Approach

1. `p[:2]` only views the first two elements.
2. `p[:]` slices the whole array, still aliasing it.

## Solution

```go
func AsSlice(p *[4]int) []int {
	return p[:]
}
```

## Walkthrough

The bug truncates to length 2. `p[:]` yields a length-4 slice backed by the same array, so writes propagate.

## Pitfalls

- `p[:]` gives all 4 elements; `p[:2]` only the first two.
- The resulting slice aliases the array's memory.
