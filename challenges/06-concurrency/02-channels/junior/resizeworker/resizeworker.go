// Package resizeworker — Gopher Workplace challenge.
package resizeworker

// ScaleRequest hands the source width to a resize worker over an unbuffered
// request channel and receives the retina (2x) width back on a reply
// channel.
//
// Examples:
//
//	ScaleRequest(640) => 1280
//	ScaleRequest(0) => 0
//	ScaleRequest(-3) => -6
func ScaleRequest(width int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
