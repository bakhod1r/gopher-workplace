# Epoll

## Solution

```go
func (e *Epoll) Wait() bool {
	return e.Active
}
```
