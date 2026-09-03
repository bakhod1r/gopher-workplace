// Package ifaceescape — Gopher Workplace challenge.
package ifaceescape

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
//	Checksum([]int{1}) => 49
func Checksum(vals []int) int {
	// CHANGE CODE BELOW THIS LINE
	var buf [64]byte
	b := buf[:0]
	for _, v := range vals {
		b = strconv.AppendInt(b, int64(v), 10)
	}
	var c counter
	w := newSink(&c)
	w.Write(b)
	return c.n
	// CHANGE CODE ABOVE THIS LINE
}
