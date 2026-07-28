// Package perms clears permission bits with AND-NOT.
package perms

// Permission is a single-bit permission flag.
type Permission uint8

// Read, Write, Execute bit flags via iota.
//
// TODO(candidate): define with 1<<iota.
const (
	Read    Permission = 0
	Write   Permission = 0
	Execute Permission = 0
)

// Revoke clears the bits in drop from set, leaving the rest intact.
//
// TODO(candidate): implement using the AND-NOT operator.
func Revoke(set, drop Permission) Permission {
	panic("not implemented")
}
