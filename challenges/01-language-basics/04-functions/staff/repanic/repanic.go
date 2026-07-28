// Package repanic recovers only a known sentinel panic and re-panics anything
// else. A planted bug swallows every panic, hiding unexpected failures.
package repanic

// ErrSentinel is the only panic value this guard is allowed to absorb.
const ErrSentinel = "handled"

// Run calls f. If f panics with ErrSentinel it returns false (absorbed). If f
// panics with anything else, that panic must propagate. If f is normal it
// returns true.
func Run(f func()) (normal bool) {
	defer func() {
		if r := recover(); r != nil {
			// CHANGE CODE BELOW THIS LINE
			normal = false
			// CHANGE CODE ABOVE THIS LINE
		}
	}()
	f()
	return true
}
