// Package takestr — Gopher Workplace challenge.
package takestr

import "unsafe"

// Take returns the first n bytes of buf as a string the caller keeps.
//
// buf is a scratch buffer the caller reuses, so the result must not be a
// view of it — a string that changes is a contradiction the rest of the
// program is not prepared for.
//
// Examples:
//
//	Take([]byte("hello"), 2) => "he", independent of buf
func Take(buf []byte, n int) string {
	// CHANGE CODE BELOW THIS LINE
	if n <= 0 || len(buf) == 0 {
		return ""
	}
	if n > len(buf) {
		n = len(buf)
	}
	return unsafe.String(unsafe.SliceData(buf), n)
	// CHANGE CODE ABOVE THIS LINE
}
