# Proxy Pattern

## Solution

```go
func (p *Proxy) Do() string {
	if p.role == "admin" {
		return p.w.Do()
	}
	return "denied"
}
```
