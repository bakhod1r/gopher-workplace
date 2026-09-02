// Package batchtotal — Gopher Workplace challenge.
package batchtotal

// BatchTotal stages the day's amounts on a buffered channel, closes it, and
// sums the channel with a range loop.
//
// Examples:
//
//	BatchTotal([]int{100, 250, 99}) => 449
//	BatchTotal(nil) => 0
//	BatchTotal([]int{-200, 200}) => 0
func BatchTotal(amounts []int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
