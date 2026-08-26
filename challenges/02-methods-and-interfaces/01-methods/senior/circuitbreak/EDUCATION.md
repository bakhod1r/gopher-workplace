# Circuit Breaker Pattern

## Solution

```go
func (b *Breaker) Call(fn func() error) error {
	if b.IsOpen {
		return errors.New("circuit open")
	}
	err := fn()
	if err != nil {
		b.ConsecutiveFails++
		if b.ConsecutiveFails >= b.Threshold {
			b.IsOpen = true
		}
	} else {
		b.ConsecutiveFails = 0
	}
	return err
}
```
