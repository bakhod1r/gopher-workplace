# Actor Router

## Solution

```go
func (r *Router) Route(msg int) {
	r.workers[r.idx].Inbox <- msg
	r.idx = (r.idx + 1) % len(r.workers)
}
```
