# Shifting to insert into a slice

## Intuition

Inserting at i needs the suffix `xs[i:]` moved into `xs[i+1:]` (rightward); copying the other way shifts left and overwrites.

## Approach

1. To open a gap at `i`, shift the tail **right**: `copy(xs[i+1:], xs[i:])`.
2. The bug copies in the wrong direction, overwriting elements.

## Solution

```go
func InsertAt(xs []int, i, v int) []int {
	xs = append(xs, 0)
	copy(xs[i+1:], xs[i:])
	xs[i] = v
	return xs
}
```

## Walkthrough

`copy(xs[i:], xs[i+1:])` shifts left and loses data. Copying `xs[i:]` into `xs[i+1:]` makes room so the new value slots in at `i`.

## Pitfalls

- Insert shifts RIGHT: `copy(xs[i+1:], xs[i:])`.
- Delete shifts LEFT: `copy(xs[i:], xs[i+1:])`.
