# Time Comparisons in Methods

## Solution

```go
func (s *Session) Ping(now time.Time) {
	s.lastPing = now
}

func (s *Session) IsExpired(now time.Time) bool {
	return now.Sub(s.lastPing) > s.timeout
}
```
