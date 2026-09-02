// Package iszerogen — Gopher Workplace challenge.
package iszerogen

// IsZero reports whether v equals the zero value of its type.
//
// Examples:
//
//	IsZero(0)   => true
//	IsZero("") => true
//	IsZero(3)   => false
func IsZero[T comparable](v T) bool {
	// TODO(candidate): compare v against the zero value of T.
	panic("not implemented")
}
