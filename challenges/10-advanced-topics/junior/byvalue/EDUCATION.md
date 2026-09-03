# Small Values Stay On The Stack

## Intuition

A sixteen-byte copy is a couple of register moves. An allocation is a trip to the allocator plus work for the collector later. For small structs the copy wins every time.

## Approach

1. Multiply the parameter's fields in place.
2. Return the parameter.

## Solution

```go
// Point is a two-dimensional integer point.
type Point struct {
	X, Y int
}

// Scale returns p with both coordinates multiplied by f.
//
// Nothing about the result outlives the call, so nothing needs the heap.
//
// Examples:
//
// 	Scale(Point{2, 3}, 2) => Point{4, 6}
func Scale(p Point, f int) Point {
	p.X *= f
	p.Y *= f
	return p
}
```

## Walkthrough

`Scale(Point{2,3}, 2)` copies two ints in, doubles them, copies two ints out. The escape analyser sees no reference leaving the frame, so the point stays in registers.

## Pitfalls

- Returning `*Point` — an allocation to avoid a two-word copy.
- Expecting the caller's point to change; take a pointer parameter if that is the intent.
