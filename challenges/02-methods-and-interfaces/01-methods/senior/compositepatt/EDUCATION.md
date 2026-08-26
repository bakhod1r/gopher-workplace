# Composite Pattern

## Solution

```go
func (f *Folder) Size() int {
	total := 0
	for _, size := range f.Files {
		total += size
	}
	for _, sub := range f.Sub {
		total += sub.Size()
	}
	return total
}
```
