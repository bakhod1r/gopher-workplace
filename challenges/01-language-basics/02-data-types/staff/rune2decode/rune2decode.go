// Package rune2decode decodes a 2-byte UTF-8 sequence into a rune by hand.
// A planted bug uses the wrong mask on the lead byte.
package rune2decode

// Decode2 combines a 2-byte UTF-8 sequence (lead 110xxxxx, cont 10xxxxxx) into
// its rune: (lead low 5 bits << 6) | (cont low 6 bits).
func Decode2(lead, cont byte) rune {
	// CHANGE CODE BELOW THIS LINE
	return rune(lead&0x0F)<<6 | rune(cont&0x3F)
	// CHANGE CODE ABOVE THIS LINE
}
