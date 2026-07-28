// Package perms models Unix-style permission bits with iota.
package perms

// Permission is a single-bit permission flag.
type Permission uint8

// Define Read, Write, Execute as consecutive power-of-two bits using iota,
// so Read=1, Write=2, Execute=4.
//
// TODO(candidate): replace the placeholders with an iota bit-flag block.
const (
	Read    Permission = 0
	Write   Permission = 0
	Execute Permission = 0
)

// Has reports whether set contains every bit in want.
//
// TODO(candidate): implement using bitwise AND.
func Has(set, want Permission) bool {
	panic("not implemented")
}
