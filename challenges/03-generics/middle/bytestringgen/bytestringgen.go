// Package bytestringgen — Gopher Workplace challenge.
package bytestringgen

// HasPrefix reports whether v starts with prefix.
func HasPrefix[T ~string | ~[]byte](v T, prefix string) bool {
	// TODO(candidate): convert to string, then test the prefix.
	panic("not implemented")
}

// Size returns the number of bytes in v.
func Size[T ~string | ~[]byte](v T) int {
	// TODO(candidate): return the length.
	panic("not implemented")
}
