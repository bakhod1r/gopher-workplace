# Interface Polymorphism

## Solution

```go
import "math"

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
```
