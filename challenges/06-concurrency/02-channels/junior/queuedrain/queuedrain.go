// Package queuedrain — Gopher Workplace challenge.
package queuedrain

// DrainQueue receives and discards every delivery attempt left on the
// channel until it is closed, returning how many were discarded.
//
// Examples:
//
//	DrainQueue(chan 1,2,3) => 3
//	DrainQueue(closed empty) => 0
//	DrainQueue(chan 9) => 1
func DrainQueue(attempts <-chan int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
