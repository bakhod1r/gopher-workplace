# ARC Cache

## Solution

```go
func (a *ARC) Access(key int) {
	if a.T2[key] {
		return
	}
	if a.T1[key] {
		delete(a.T1, key)
		a.T2[key] = true
		return
	}
	a.T1[key] = true
}
```
