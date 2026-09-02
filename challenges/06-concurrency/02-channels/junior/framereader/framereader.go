// Package framereader — Gopher Workplace challenge.
package framereader

// ReadFrames buffers the frame sizes on a channel, closes it, then performs
// len(sizes)+extra receives and returns everything received.
//
// Receives past the end of a closed channel yield the zero size 0.
//
// Examples:
//
//	ReadFrames([]int{1024, 512}, 1) => [1024 512 0]
//	ReadFrames(nil, 2)              => [0 0]
//	ReadFrames([]int{5}, 0)         => [5]
func ReadFrames(sizes []int, extra int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
