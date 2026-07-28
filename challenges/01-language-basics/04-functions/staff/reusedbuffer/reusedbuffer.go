// Package reusedbuffer builds a reader that fills a REUSED internal buffer each
// call and returns it. A planted bug returns the buffer directly; because the
// next call resets it with buf[:0] and re-appends into the SAME backing array,
// a previously returned result is overwritten.
package reusedbuffer

// Reader returns a function that, given values, returns them as a slice. Results
// from earlier calls must remain valid after later calls.
func Reader() func(vals ...int) []int {
	buf := make([]int, 0, 16)
	return func(vals ...int) []int {
		buf = buf[:0]
		buf = append(buf, vals...)
		// CHANGE CODE BELOW THIS LINE
		return buf
		// CHANGE CODE ABOVE THIS LINE
	}
}
