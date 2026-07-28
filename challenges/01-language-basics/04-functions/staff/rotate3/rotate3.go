// Package rotate3 rotates three variables left: (a,b,c) -> (b,c,a). A planted
// bug uses sequential assignments that clobber values before they are moved.
package rotate3

// RotateLeft returns a,b,c rotated left by one position.
func RotateLeft(a, b, c int) (int, int, int) {
	// CHANGE CODE BELOW THIS LINE
	a = b
	b = c
	c = a
	// CHANGE CODE ABOVE THIS LINE
	return a, b, c
}
