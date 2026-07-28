// Package maskpanic recovers the ORIGINAL panic value. A planted bug has a
// second deferred function panic during unwind, replacing the original value
// that the outer recover then sees.
package maskpanic

// FirstPanic runs f (which panics), and returns the value of the ORIGINAL panic.
// A cleanup step must not overwrite that value by panicking itself.
func FirstPanic(f func()) (got any) {
	defer func() {
		got = recover()
	}()
	// CHANGE CODE BELOW THIS LINE
	defer func() {
		panic("cleanup-failure")
	}()
	// CHANGE CODE ABOVE THIS LINE
	f()
	return
}
