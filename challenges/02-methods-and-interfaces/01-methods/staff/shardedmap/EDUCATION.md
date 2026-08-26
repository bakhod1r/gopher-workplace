# Sharded Map

## Solution

```go
func (m *ShardedMap) Set(key string, val int) {
	s := m.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}
```
