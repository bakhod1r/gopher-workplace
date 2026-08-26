# GC Sweep

## Solution

```go
func (h *Heap) Sweep() int {
	c := 0
	for _, b := range h.Objects {
		if !b { c++ }
	}
	return c
}
```
