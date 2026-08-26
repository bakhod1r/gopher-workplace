# Reactive Streams

## Solution

```go
func (s *Stream) Filter(fn func(int) bool) *Stream {
	var res []int
	for _, v := range s.Data {
		if fn(v) { res = append(res, v) }
	}
	s.Data = res
	return s
}

func (s *Stream) Map(fn func(int) int) *Stream {
	for i, v := range s.Data {
		s.Data[i] = fn(v)
	}
	return s
}
```
