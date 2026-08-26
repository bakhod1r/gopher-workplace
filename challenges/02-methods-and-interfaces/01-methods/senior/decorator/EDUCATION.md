# Decorator Pattern

## Solution

```go
func (d *Decorator) DoWork() string {
	return "[" + d.Comp.DoWork() + "]"
}
```
