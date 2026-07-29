# Arrays as map keys

## Intuition

Unlike slices, **arrays are comparable**, so a `[2]int` (or `[N]T` of comparable
T) is a valid map key — perfect for coordinate pairs:

```go
m[[2]int{c[0], c[1]}]++
```

The array's element order *is* its identity; swapping fields makes a different
key.

## Approach

1. Bug: `[2]int{c[1], c[0]}` swaps the coordinate order in the array key.
2. Fix: `[2]int{c[0], c[1]}` keeps (row, col) order.

## Solution

```go
func CountCells(cells [][2]int) map[[2]int]int {
	m := make(map[[2]int]int)
	for _, c := range cells {
		m[[2]int{c[0], c[1]}]++
	}
	return m
}
```

## Walkthrough

cell [0,1]: buggy key [1,0], correct key [0,1]. Two [0,1] cells -> [0,1]:2; [2,3] -> 1.

## Pitfalls

- Slices, maps, and funcs are **not** comparable — can't be keys.
- A struct of comparable fields is also a valid key.
- Array keys are compared and hashed by value (all elements).
