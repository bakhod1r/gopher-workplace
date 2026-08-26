# Bloom Filter

## Solution

```go
func (f *Filter) Add(item string) {
	if len(item) == 0 { return }
	f.bits[hash1(item)] = true
	f.bits[hash2(item)] = true
}

func (f *Filter) MightContain(item string) bool {
	if len(item) == 0 { return false }
	return f.bits[hash1(item)] && f.bits[hash2(item)]
}
```
