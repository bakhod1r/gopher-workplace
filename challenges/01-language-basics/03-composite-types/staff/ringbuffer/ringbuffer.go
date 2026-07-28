// Package ringbuffer indexes a circular buffer. A planted bug omits the modulo,
// so logical indices past the end go out of range.
package ringbuffer

// At returns the element at logical position i in a ring buffer whose physical
// storage is buf and whose logical start is head. i may be >= len(buf).
func At(buf []int, head, i int) int {
	// CHANGE CODE BELOW THIS LINE
	return buf[head+i]
	// CHANGE CODE ABOVE THIS LINE
}
