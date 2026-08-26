# Thread-Safe Methods

## Intuition

Methods that mutate shared state must use the embedded mutex. `l.mu.Lock()` and
`defer l.mu.Unlock()` is the idiomatic way to protect struct fields.

## Solution

```go
func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.tokens > 0 {
		l.tokens--
		return true
	}
	return false
}

func (l *Limiter) Refill(n int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.tokens += n
}
```
