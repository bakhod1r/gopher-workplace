// Package repeatbytes — Gopher Workplace challenge.
package repeatbytes

// Repeat returns n concatenated copies of b in a freshly allocated
// slice that shares nothing with b.
//
// For n <= 0 the result is empty. The allocation must happen once, at the
// final size.
//
// Examples:
//
//	Repeat([]byte("ab"), 2) => []byte("abab")
func Repeat(b []byte, n int) []byte {
	panic("not implemented")
}
