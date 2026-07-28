// Package varint decodes an unsigned LEB128 varint (as used by protobuf).
// A planted bug advances the shift by 8 bits instead of 7.
package varint

// Decode reads a base-128 varint from b, returning the value and the number of
// bytes consumed. Each byte contributes 7 bits; the high bit means "more".
// Returns (0,0) if the slice ends mid-varint.
func Decode(b []byte) (uint64, int) {
	var v uint64
	var shift uint
	for i := 0; i < len(b); i++ {
		c := b[i]
		v |= uint64(c&0x7F) << shift
		if c&0x80 == 0 {
			return v, i + 1
		}
		// CHANGE CODE BELOW THIS LINE
		shift += 8
		// CHANGE CODE ABOVE THIS LINE
	}
	return 0, 0
}
