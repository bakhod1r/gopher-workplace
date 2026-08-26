# Concurrent Map

## Solution

```go
func (m *ConcurrentMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[key]
	return val, ok
}

func (m *ConcurrentMap) Set(key string, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}
```
