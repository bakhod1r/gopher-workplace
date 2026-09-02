// Package searchresults — Gopher Workplace challenge.
package searchresults

// TopResults takes the first n hits off the search stream and then drains the
// rest so the producer goroutine can finish instead of leaking. n <= 0 takes
// nothing.
//
// Examples:
//
//	TopResults(stream of 5 hits, 3)  => first 3 hits
//	TopResults(stream of 2 hits, 9)  => both hits
//	TopResults(stream, 0)            => nil
func TopResults(hits <-chan string, n int) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
