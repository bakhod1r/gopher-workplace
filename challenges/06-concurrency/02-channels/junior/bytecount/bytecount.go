// Package bytecount — Gopher Workplace challenge.
package bytecount

// TotalBytes receives every response size from sizes until the channel is
// closed and returns the total number of bytes served.
//
// Examples:
//
//	TotalBytes(chan 1200, 800) => 2000
//	TotalBytes(chan 512)       => 512
//	TotalBytes(closed empty)   => 0
func TotalBytes(sizes <-chan int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
