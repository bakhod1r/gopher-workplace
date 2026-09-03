// Package bigscratch — Gopher Workplace challenge.
package bigscratch

import "strconv"

// Format returns v's decimal digits.
//
// The result escapes, so whatever it points into escapes with it — sizing
// the scratch buffer for the worst imaginable case then costs that much
// heap on every call.
//
// Examples:
//
//	Format(42) => []byte("42")
func Format(v int64) []byte {
	// CHANGE CODE BELOW THIS LINE
	var buf [4096]byte
	return strconv.AppendInt(buf[:0], v, 10)
	// CHANGE CODE ABOVE THIS LINE
}
