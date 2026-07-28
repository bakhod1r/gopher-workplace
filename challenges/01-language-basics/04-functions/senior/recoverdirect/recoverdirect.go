// Package recoverdirect guards a call that may panic. A planted bug calls
// recover() directly in the body instead of from a deferred function, so it
// never actually stops the panic.
package recoverdirect

// Guard runs f and returns true if f panicked (and was recovered), false if it
// returned normally.
func Guard(f func()) (ok bool) {
	// CHANGE CODE BELOW THIS LINE
	if r := recover(); r != nil {
		ok = true
	}
	f()
	// CHANGE CODE ABOVE THIS LINE
	return ok
}
