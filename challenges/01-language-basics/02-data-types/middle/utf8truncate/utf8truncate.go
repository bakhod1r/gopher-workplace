// Package utf8truncate truncates a string to a byte budget without splitting a
// rune.
package utf8truncate

// Truncate returns the longest prefix of s whose length in bytes is <= max,
// never cutting through the middle of a multi-byte rune.
//
// TODO(candidate): implement by ranging (byte index + rune width).
func Truncate(s string, max int) string {
	panic("not implemented")
}
