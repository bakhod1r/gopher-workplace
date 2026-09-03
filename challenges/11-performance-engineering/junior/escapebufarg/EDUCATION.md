# Returning A Value Costs Nothing

## Intuition

The compiler heap-allocates only what it cannot prove stays local. Returning a value proves nothing escapes; returning a pointer proves the opposite.

## Approach

1. `Add` builds and returns a `Point` literal.
2. `AddInto` assigns a `Point` literal through the pointer.

## Solution

```go
func Add(a, b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func AddInto(dst *Point, a, b Point) {
	*dst = Point{a.X + b.X, a.Y + b.Y}
}
```

## Walkthrough

`AddInto(&p, p, ...)` works because `a` is a copy taken at call time, so overwriting `*dst` cannot disturb the operands mid-computation.

## Pitfalls

- Returning `*Point` from `Add`, which forces the value onto the heap.
- Writing the fields one at a time in `AddInto` — correct here, but it breaks the moment `dst` aliases an operand you still need.
- Assuming the escape happens because of the pointer *type*; it happens because the address outlives the frame.
