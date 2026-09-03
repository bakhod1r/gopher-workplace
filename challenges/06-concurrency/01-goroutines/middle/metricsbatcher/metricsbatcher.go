// Package metricsbatcher — Gopher Workplace challenge.
package metricsbatcher

// FlushBatches cuts points into batches of batchSize, flushes each batch in its
// own goroutine, and reports how many points the collector accepted plus the
// points that must be re-queued, in their original order. A rejected batch is
// retried whole, so its points come back untouched.
//
// A batchSize of zero or less flushes nothing and re-queues everything.
//
// Examples:
//
//	FlushBatches([]int{1, 2, 3, 4}, 2, flush)  => 4, []
//	FlushBatches([]int{1, -1, 3}, 2, flush)    => 1, [1 -1]
//	FlushBatches(nil, 2, flush)                => 0, []
func FlushBatches(points []int, batchSize int, flush func(batch []int) error) (int, []int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
