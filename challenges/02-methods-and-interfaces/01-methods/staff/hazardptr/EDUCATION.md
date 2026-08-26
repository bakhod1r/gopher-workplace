# Hazard Pointers

## Solution

```go
func (h *Hazard) Protect(shared *atomic.Pointer[int]) *int {
	p := shared.Load()
	h.ptr.Store(p)
	if shared.Load() == p {
		return p
	}
	return nil
}
```
