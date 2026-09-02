// Package telemetryflusher — Gopher Workplace challenge.
package telemetryflusher

// BatchErrorCounts returns how many server errors each flush batch contains.
//
// Examples:
//
//	BatchErrorCounts([]int{200, 500, 503, 200}, 2)  => [1 1]
//	BatchErrorCounts([]int{500, 500, 200}, 2)       => [2 0]
//	BatchErrorCounts([]int{200}, 0)                 => []
func BatchErrorCounts(codes []int, batch int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
