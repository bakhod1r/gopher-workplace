// Package freezebytes — Gopher Workplace challenge.
package freezebytes

import "unsafe"

// Snapshot returns a byte slice holding s's bytes that the caller may
// modify freely.
//
// A view over the string's own bytes is not that: strings may live in
// read-only memory, and every other holder of s would see the writes.
//
// Examples:
//
//	b := Snapshot("hi"); b[0] = 'H' => s is unaffected
func Snapshot(s string) []byte {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
	// CHANGE CODE ABOVE THIS LINE
}
