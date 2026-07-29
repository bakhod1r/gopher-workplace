# Append aliasing through shared capacity

## Intuition

Two appends to a base with spare capacity write the same memory; a three-index slice (`s[:len:len]`) caps it so the next append copies.

## Approach

1. Two appends sharing spare capacity write to the same array.
2. Clip the base: `base[:len(base):len(base)]` so each append reallocates.

## Solution

```go
func Fork(base []int, a, b int) ([]int, int) {
	bc := base[:len(base):len(base)]
	x := append(bc, a)
	y := append(bc, b)
	_ = y
	return x, x[len(x)-1]
}
```

## Walkthrough

Without clipping, appending `b` overwrites the slot `x` holds, so `x[1]` becomes 200. The full-slice clip isolates the two forks.

## Pitfalls

- Appending to a base with spare cap twice makes the results alias.
- `base[:len:len]` guarantees the next append reallocates.
