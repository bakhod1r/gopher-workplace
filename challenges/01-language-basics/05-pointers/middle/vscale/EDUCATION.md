# Value receivers and immutability

## Intuition

A value-receiver method works on a copy; returning a new value keeps the original immutable — natural for small value types.

## Approach

1. A value receiver is fine: we return a new `Point`, not mutate.
2. Multiply each field by `k` and construct the result.

## Solution

```go
type Point struct{ X, Y int }

func (p Point) Scaled(k int) Point {
	return Point{X: p.X * k, Y: p.Y * k}
}
```

## Walkthrough

`Point{1,2}.Scaled(3)` returns `Point{3, 6}`; the original is unchanged because the receiver is a copy.

## Pitfalls

- Mutating `p` inside a value receiver changes only the copy.
- Return a new struct to express a transform.
