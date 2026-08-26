# HyperLogLog

## Solution

```go
func (h *HLL) Add(hash uint32) {
	zeros := leadingZeros(hash)
	if zeros > h.maxZeros {
		h.maxZeros = zeros
	}
}
```
