// Package bufconcat concatenates two byte slices into a preallocated buffer.
// A planted bug copies b over the start instead of after a.
package bufconcat

// Concat returns a new buffer containing a followed by b.
func Concat(a, b []byte) []byte {
	out := make([]byte, len(a)+len(b))
	copy(out, a)
	// CHANGE CODE BELOW THIS LINE
	copy(out, b)
	// CHANGE CODE ABOVE THIS LINE
	return out
}
