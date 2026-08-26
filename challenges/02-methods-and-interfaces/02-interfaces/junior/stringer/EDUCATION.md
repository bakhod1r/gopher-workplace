# fmt.Stringer

## Solution

```go
func (c Color) String() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
```
