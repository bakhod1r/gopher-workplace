// Package perms clears permission bits. A planted XOR bug toggles instead.
package perms

// Permission is a single-bit permission flag.
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Execute
)

// Revoke clears the bits in drop from set. Dropping an absent bit is a no-op.
func Revoke(set, drop Permission) Permission {
	// CHANGE CODE BELOW THIS LINE
	return set ^ drop
	// CHANGE CODE ABOVE THIS LINE
}
