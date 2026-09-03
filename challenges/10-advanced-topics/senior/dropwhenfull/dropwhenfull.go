// Package dropwhenfull — Gopher Workplace challenge.
package dropwhenfull

// Offer sends v on ch if it can be accepted immediately, and reports
// whether it was.
//
// A metrics pipeline must never block its caller: when the buffer is full,
// the sample is dropped.
//
// Examples:
//
//	Offer(ch, 1) => true when ch has room
func Offer(ch chan<- int, v int) bool {
	panic("not implemented")
}
