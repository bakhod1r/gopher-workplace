// Package hexencode encodes bytes as a lowercase hex string.
package hexencode

// Encode returns the lowercase hex encoding of b (two chars per byte).
//
// TODO(candidate): implement without encoding/hex; map each nibble to 0-9a-f.
func Encode(b []byte) string {
	panic("not implemented")
}
