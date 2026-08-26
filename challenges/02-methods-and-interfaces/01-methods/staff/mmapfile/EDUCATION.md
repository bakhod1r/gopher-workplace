# Mmap

## Solution

```go
func (m *Mmap) ReadByteAt(pos int) byte {
	return m.Data[pos]
}
```
