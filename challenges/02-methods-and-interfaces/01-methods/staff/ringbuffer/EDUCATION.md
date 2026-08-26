# Ring Buffer

## Solution

```go
func (r *RingBuffer) Push(val int) error {
	if r.size == len(r.data) {
		return errors.New("full")
	}
	r.data[r.tail] = val
	r.tail = (r.tail + 1) % len(r.data)
	r.size++
	return nil
}
```
