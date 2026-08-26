# Adapter Pattern

## Solution

```go
func (a *ModernAdapter) GetIntData() int {
	s := a.legacy.GetStringData()
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
```
