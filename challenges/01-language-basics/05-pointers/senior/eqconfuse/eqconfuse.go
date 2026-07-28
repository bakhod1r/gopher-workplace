// Package eqconfuse checks whether two pointers refer to the same variable. A
// planted bug compares the pointed-to VALUES, so two distinct variables holding
// equal values are wrongly reported as the same.
package eqconfuse

// Same reports whether a and b point to the SAME int (same address).
func Same(a, b *int) bool {
	// CHANGE CODE BELOW THIS LINE
	return *a == *b
	// CHANGE CODE ABOVE THIS LINE
}
