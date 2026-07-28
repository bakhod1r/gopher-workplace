// Package valuerecv increments a counter via a method. A planted bug uses a
// VALUE receiver, so the increment mutates a copy and is lost.
package valuerecv

type Counter struct{ N int }

// Inc should increment the counter so the change persists on the caller.
// CHANGE CODE BELOW THIS LINE
func (c Counter) Inc() {
	// CHANGE CODE ABOVE THIS LINE
	c.N++
}
