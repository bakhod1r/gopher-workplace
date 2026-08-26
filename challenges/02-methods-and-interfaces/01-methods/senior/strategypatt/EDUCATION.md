# Strategy Pattern (Functional)

## Solution

```go
func (c *Context) Process(strategy func(int) int) {
	for i, v := range c.Data {
		c.Data[i] = strategy(v)
	}
}
```
