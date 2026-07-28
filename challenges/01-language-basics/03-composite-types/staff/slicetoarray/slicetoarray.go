// Package slicetoarray reads a fixed 4-byte header. A planted bug converts
// without checking length, so short input panics.
package slicetoarray

// First4 returns the first 4 bytes of b as an array, and ok=false if b is too
// short (it must not panic).
func First4(b []byte) ([4]byte, bool) {
	// CHANGE CODE BELOW THIS LINE
	return [4]byte(b[:4]), true
	// CHANGE CODE ABOVE THIS LINE
}
