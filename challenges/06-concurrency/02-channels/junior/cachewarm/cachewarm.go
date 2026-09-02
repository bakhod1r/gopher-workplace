// Package cachewarm — Gopher Workplace challenge.
package cachewarm

// NextKey performs one receive on the eviction feed and reports whether a
// real key id arrived. A closed, drained feed yields 0, false.
//
// Examples:
//
//	NextKey(chan 5)       => 5, true
//	NextKey(closed empty) => 0, false
//	NextKey(chan 0)       => 0, true
func NextKey(feed <-chan int) (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
