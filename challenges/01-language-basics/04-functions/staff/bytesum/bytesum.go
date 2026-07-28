// Package bytesum sums byte values. A planted bug accumulates into a uint8,
// which overflows (wraps mod 256) for totals above 255.
package bytesum

// Sum returns the total of the byte values in bs.
func Sum(bs []byte) int {
	// CHANGE CODE BELOW THIS LINE
	var total uint8
	for _, b := range bs {
		total += b
	}
	return int(total)
	// CHANGE CODE ABOVE THIS LINE
}
