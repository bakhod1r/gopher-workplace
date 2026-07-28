// Package bareoverwrite doubles a result via a deferred adjust. A planted bug
// never assigns the named return before a bare `return`, so the deferred
// doubling acts on zero.
package bareoverwrite

// Doubled computes x*2 by assigning the named return then letting a deferred
// closure double it. A planted bug forgets to assign result, so it returns 0.
func Doubled(x int) (result int) {
	defer func() { result *= 2 }()
	local := x
	_ = local
	// CHANGE CODE BELOW THIS LINE
	return
	// CHANGE CODE ABOVE THIS LINE
}
