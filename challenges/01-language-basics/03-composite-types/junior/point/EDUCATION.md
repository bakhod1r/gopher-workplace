# Structs

## Intuition

A struct groups named fields into one value:

```go
type Point struct { X, Y int }
p := Point{X: 1, Y: 2}
```

Fields are accessed with dot notation; the whole struct copies on assignment.

## Approach

1. Build a brand-new Point with X=p.X+dx and Y=p.Y+dy.
2. Return it. Because Point is a value and we return a new one, the caller's p is unchanged.

## Solution

```go
type Point struct {
	X int
	Y int
}

func Translate(p Point, dx, dy int) Point {
	return Point{X: p.X + dx, Y: p.Y + dy}
}
```

## Walkthrough

Translate(Point{1,2},3,4): construct Point{1+3, 2+4} = Point{4,6}; the argument copy is discarded, caller keeps {1,2}.

## Pitfalls

- Struct assignment copies all fields (value semantics).
- Use a pointer receiver / pointer to mutate a struct in a function.
- Field order affects memory layout (padding), not equality.
