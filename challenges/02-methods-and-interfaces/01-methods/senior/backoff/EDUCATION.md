# Exponential Backoff

## Solution

```go
func (b *Backoff) Next() time.Duration {
	ret := b.current
	b.current *= 2
	if b.current > b.max {
		b.current = b.max
	}
	return ret
}
```
