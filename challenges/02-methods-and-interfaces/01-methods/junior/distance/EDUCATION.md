# Methods with Parameters

## Intuition

A method can take parameters *in addition to* its receiver. `p.DistanceTo(other)`
has one receiver (`p`) and one explicit parameter (`other`) — both are `Point`.

## Approach

1. Compute `dx = other.X - p.X` and `dy = other.Y - p.Y`.
2. Return `math.Sqrt(dx*dx + dy*dy)`.

## Solution

```go
func (p Point) DistanceTo(other Point) float64 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	return math.Sqrt(dx*dx + dy*dy)
}
```

## Walkthrough

For `Point{0,0}` → `Point{3,4}`:
- `dx` = 3, `dy` = 4.
- `dx*dx + dy*dy` = 9 + 16 = 25.
- `√25` = 5.

## Pitfalls

- Using `math.Pow(dx, 2)` instead of `dx*dx` — works but slower.
- Forgetting to import `"math"`.
