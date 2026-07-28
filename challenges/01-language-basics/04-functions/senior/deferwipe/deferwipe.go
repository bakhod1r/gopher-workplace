// Package deferwipe should return a computed value, but a planted deferred
// closure resets the named return to zero, wiping the result.
package deferwipe

// Compute returns a*b, but a stray deferred closure zeroes the result.
func Compute(a, b int) (result int) {
	// CHANGE CODE BELOW THIS LINE
	defer func() { result = 0 }()
	// CHANGE CODE ABOVE THIS LINE
	result = a * b
	return
}
