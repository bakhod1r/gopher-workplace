# Count-Min Sketch

## Solution

```go
func (s *Sketch) Add(item string) {
	if len(item) == 0 { return }
	s.row1[h1(item)]++
	s.row2[h2(item)]++
}

func (s *Sketch) Count(item string) int {
	if len(item) == 0 { return 0 }
	c1 := s.row1[h1(item)]
	c2 := s.row2[h2(item)]
	if c1 < c2 { return c1 }
	return c2
}
```
