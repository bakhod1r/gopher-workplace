# The Buffer That An Interface Sent To The Heap

## Intuition

Escape analysis works on what the compiler can prove. Handing a slice to an interface method destroys the proof: the dynamic implementation is not known at compile time, so the argument must be assumed to escape.

## Approach

1. Render into the local buffer as before.
2. Sum the bytes with a direct loop instead of routing them through the interface.

## Solution

```go
import "strconv"

// sink accepts the rendering; it is a stand-in for a real writer.
type sink interface{ Write(p []byte) (int, error) }

type counter struct{ n int }

func (c *counter) Write(p []byte) (int, error) {
	for _, b := range p {
		c.n += int(b)
	}
	return len(p), nil
}

// newSink is a variable, so the compiler cannot devirtualise the interface
// value it produces.
var newSink = func(c *counter) sink { return c }

// Checksum renders vals as decimal digits into a scratch buffer and
// returns the sum of the bytes written.
//
// Passing the scratch buffer to an interface makes it escape. Everything
// here can stay in the frame.
//
// Examples:
//
// 	Checksum([]int{1}) => 49
func Checksum(vals []int) int {
	var buf [64]byte
	b := buf[:0]
	for _, v := range vals {
		b = strconv.AppendInt(b, int64(v), 10)
	}
	n := 0
	for _, x := range b {
		n += int(x)
	}
	return n
}
```

## Walkthrough

`w.Write(b)` makes both `buf` and `c` escape — two allocations per call. Summing `b` inline keeps the 64-byte array in the frame and reports 0 allocations.

## Pitfalls

- Calling `c.Write(b)` on the concrete type instead — better, but the fixture still exists to be removed.
- Assuming a small local array can never escape; its address leaving the function is what matters, not its size.
