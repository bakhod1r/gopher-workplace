# Observer Pattern

## Solution

```go
func (s *Subject) SetState(val int) {
	s.state = val
	for _, o := range s.observers {
		o(val)
	}
}
```
