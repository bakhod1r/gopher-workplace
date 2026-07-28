// Package recovertype recovers a panic and returns it as an error, only when the
// panic value IS an error. A planted bug asserts the wrong type, dropping error
// panics.
package recovertype

// Call runs f. If f panics with an error value, Call returns it. If f panics
// with a non-error, Call returns nil. Normal completion returns nil.
func Call(f func()) (err error) {
	defer func() {
		r := recover()
		// CHANGE CODE BELOW THIS LINE
		if s, ok := r.(string); ok {
			_ = s
		}
		// CHANGE CODE ABOVE THIS LINE
	}()
	f()
	return
}
