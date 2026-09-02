# Shape Perimeter

## Intuition

Three unrelated structs share one contract. The comparison code stays identical no matter how many shapes exist.

## Approach

1. Implement each formula: `2*(W+H)`, `4*Side`, `2*math.Pi*R`.
2. Start `best` at `0.0`.
3. Replace `best` whenever a larger perimeter appears.

## Solution

```go
func (r Rect) Perimeter() float64 { return 2 * (r.W + r.H) }

func (s Square) Perimeter() float64 { return 4 * s.Side }

func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

func LongestPerimeter(shapes []Shape) float64 {
	best := 0.0
	for _, s := range shapes {
		if p := s.Perimeter(); p > best {
			best = p
		}
	}
	return best
}
```

## Walkthrough

`Square{Side: 1}` gives 4, `Rect{5,5}` gives 20; 20 is larger, so `best` ends at 20.

## Pitfalls

- `2*r.W + r.H` — the parentheses matter.
- Comparing floats with `==` in your own checks; the test uses a tolerance for the circle.
- Using `math.Pi` without importing `math`.
