// Package runlength run-length-encodes a byte string.
package runlength

// Encode returns the run-length encoding of s as pairs: for each run, the
// character followed by its count as decimal. E.g. "aaab" -> "a3b1".
//
// TODO(candidate): scan runs, emit char + count.
func Encode(s string) string {
	panic("not implemented")
}
