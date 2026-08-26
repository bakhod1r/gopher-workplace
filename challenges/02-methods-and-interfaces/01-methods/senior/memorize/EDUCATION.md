# Memoization

## Solution

```go
func (m *Memoizer) Get(key string) string {
	if val, ok := m.cache[key]; ok {
		return val
	}
	val := m.fn(key)
	m.cache[key] = val
	return val
}
```
