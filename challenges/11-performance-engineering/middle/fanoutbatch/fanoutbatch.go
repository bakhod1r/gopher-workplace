// Package fanoutbatch — Gopher Workplace challenge.
package fanoutbatch

// Chunks splits items into ceil(len/size) contiguous batches, the last one
// possibly short. The batches are sub-slices of the input, so they share its
// memory. A non-positive size, or no items, gives an empty, non-nil result.
//
// Examples:
//
//	Chunks([]int{1, 2, 3, 4, 5}, 2) => [][]int{{1,2},{3,4},{5}}
func Chunks(items []int, size int) [][]int {
	panic("not implemented")
}

// SumBatches processes the batches concurrently — one goroutine per batch,
// which is the point of batching: the fan-out width is the batch count, not
// the item count — and returns each batch's sum in batch order.
//
// Examples:
//
//	SumBatches([]int{1, 2, 3, 4, 5}, 2) => []int{3, 7, 5}
func SumBatches(items []int, size int) []int {
	panic("not implemented")
}
