// Package metricsflush — Gopher Workplace challenge.
package metricsflush

// FlushBatches reads samples from in and sends them onward as slices of at
// most size elements. A batch is emitted as soon as it is full; when in is
// closed, any partial batch is emitted too. The returned channel is closed
// once in is drained. A size of zero or less means one sample per batch.
//
// Examples:
//
//	in: 1,2,3 size 2  => [1 2] [3]
//	in: 1,2   size 5  => [1 2]
//	in: none  size 3  => (no batches)
func FlushBatches(in <-chan int, size int) <-chan []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
