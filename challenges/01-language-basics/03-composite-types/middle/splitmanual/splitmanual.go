// Package splitmanual splits a string on a single-byte separator by hand.
package splitmanual

// Split returns the fields of s separated by sep (a single byte), without using
// strings.Split. Consecutive separators yield empty fields; result always has
// (number of seps)+1 elements.
//
// TODO(candidate): scan, cutting at each sep.
func Split(s string, sep byte) []string {
	panic("not implemented")
}
