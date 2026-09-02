// Package pagefeed — Gopher Workplace challenge.
package pagefeed

// StreamPages runs a producer goroutine that emits page numbers 1..n on an
// unbuffered channel and returns the page numbers the exporter collected.
//
// Examples:
//
//	StreamPages(3) => [1 2 3]
//	StreamPages(1) => [1]
//	StreamPages(0) => []
func StreamPages(n int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
