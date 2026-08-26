# Flux Pattern

## Solution

```go
func (s *Store) Dispatch(action string) {
	switch action {
	case "INC": s.Count++
	case "DEC": s.Count--
	}
}
```
