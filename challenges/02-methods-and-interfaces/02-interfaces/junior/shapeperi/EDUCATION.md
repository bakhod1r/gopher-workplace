# Shape Perimeter

## Solution

```go
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
func (s Square) Perimeter() float64 { return 4 * s.Side }
```
