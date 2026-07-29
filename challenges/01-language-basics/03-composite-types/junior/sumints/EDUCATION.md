# Ranging over slices

## Intuition

`for i, v := range xs` visits each element; use `_` for an index you don't need.
Accumulate into a variable declared before the loop:

```go
total := 0
for _, x := range xs { total += x }
```

## Approach

1. Initialize total to 0.
2. Range over xs, adding each element to total.
3. Return total; empty and nil slices skip the loop and yield 0.

## Solution

```go
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
```

## Walkthrough

Sum([-5,5]): total 0 -> -5 -> 0. Return 0.

## Pitfalls

- A `nil` slice is safe to range (zero iterations); no need to check.
- The range value is a **copy**; assigning to it doesn't change the slice.
- `len(xs)` is O(1); indexing `xs[i]` is fine too, but range reads cleaner.
