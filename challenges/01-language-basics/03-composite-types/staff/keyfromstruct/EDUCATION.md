# Structs as map keys

## Intuition

A struct whose fields are all comparable is itself comparable and can be a map
key. Its identity is the tuple of field values, so order/assignment matters:

```go
m[Point{p.X, p.Y}]++ // or simply m[p]++
```

## Approach

1. Bug: `Point{p.Y, p.X}` swaps fields, so (1,2) is tallied under key (2,1).
2. Fix: `Point{p.X, p.Y}` preserves the coordinate order.

## Solution

```go
type Point struct {
	X, Y int
}

func Count(pts []Point) map[Point]int {
	m := make(map[Point]int)
	for _, p := range pts {
		m[Point{p.X, p.Y}]++
	}
	return m
}
```

## Walkthrough

For p={1,2}: buggy key {2,1}, correct key {1,2}. Two identical {1,2} inputs both map to {1,2} -> count 2; {3,4} -> count 1.

## Pitfalls

- A struct with a slice/map/func field is **not** comparable — compile error as a
  key.
- `m[p]` reuses the loop value directly, avoiding field mistakes.
- Keys are hashed/compared by all fields.
