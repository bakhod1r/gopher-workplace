# Lock-Free Stack (CAS)

## Solution

```go
func (s *Stack) Push(val int) {
	n := &node{val: val}
	for {
		old := s.head.Load()
		n.next = old
		if s.head.CompareAndSwap(old, n) {
			break
		}
	}
}
```
