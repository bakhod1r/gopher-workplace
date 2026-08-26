# Chain of Responsibility

## Solution

```go
func (h *H20) Handle(req int) string {
	if req == 20 {
		return "twenty"
	}
	return h.Next(req)
}
```
