# RCU Pattern

## Solution

```go
func (r *RCU) Update(newData string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ptr.Store(&Config{Data: newData})
}
```
