// Package metricdrop — Gopher Workplace challenge.
package metricdrop

// TryRecord attempts one non-blocking send of sample on the metrics buffer
// using select with default. It reports whether the sample was recorded.
//
// A full buffer yields false rather than blocking the caller.
//
// Examples:
//
//	TryRecord(buffer cap 1 empty, 5) => true
//	TryRecord(buffer cap 1 full, 5)  => false
//	TryRecord(unbuffered buffer, 5)  => false
func TryRecord(buffer chan<- int, sample int) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
