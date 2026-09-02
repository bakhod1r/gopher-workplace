# Shape Area

## Intuition

`TotalArea` depends on the *ability to be measured*, not on any shape. New shapes plug in for free.

## Approach

1. `Rect.Area` returns `r.W * r.H`.
2. `Square.Area` returns `s.Side * s.Side`.
3. `TotalArea` accumulates `s.Area()` into a `float64`.

## Solution

```go
func (r Rect) Area() float64 { return r.W * r.H }

func (s Square) Area() float64 { return s.Side * s.Side }

func TotalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}
```

## Walkthrough

`{Rect{2,2}, Square{3}}`: 4 + 9 = 13.

## Pitfalls

- Starting the accumulator as `0` (an int) and getting a type error on `+=`.
- Using `math.Pow(s.Side, 2)` — correct but slower and needs an import.
- Returning after the first shape instead of summing.
