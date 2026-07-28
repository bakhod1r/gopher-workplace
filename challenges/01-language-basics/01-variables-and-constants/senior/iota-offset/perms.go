// Package perms packs permission bits. One planted bug shifts every flag.
package perms

// Permission is a single-bit permission flag.
type Permission uint8

const (
	// CHANGE CODE BELOW THIS LINE
	Read Permission = 1 << (iota + 1)
	// CHANGE CODE ABOVE THIS LINE
	Write
	Execute
)

// Has reports whether set contains every bit in want.
func Has(set, want Permission) bool { return set&want == want }
