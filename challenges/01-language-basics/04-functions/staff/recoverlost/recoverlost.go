// Package recoverlost converts a panic into a returned error. A planted bug
// assigns the recovered value to a LOCAL variable instead of the named return,
// so the caller receives nil.
package recoverlost

import "fmt"

// Safe runs f; if f panics, Safe returns an error describing it, else nil.
func Safe(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// CHANGE CODE BELOW THIS LINE
			e := fmt.Errorf("panic: %v", r)
			_ = e
			// CHANGE CODE ABOVE THIS LINE
		}
	}()
	f()
	return nil
}
