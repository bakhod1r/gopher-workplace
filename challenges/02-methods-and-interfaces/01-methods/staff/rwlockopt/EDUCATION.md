# Lock Escalation

## Solution

```go
func (o *OptLock) IncrementIfZero() int {
	o.mu.RLock()
	v := o.v
	o.mu.RUnlock()

	if v != 0 {
		return v
	}

	o.mu.Lock()
	defer o.mu.Unlock()
	if o.v == 0 {
		o.v++
	}
	return o.v
}
```
