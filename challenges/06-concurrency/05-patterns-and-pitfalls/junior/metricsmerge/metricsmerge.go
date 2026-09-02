// Package metricsmerge — Gopher Workplace challenge.
package metricsmerge

// MergeMetrics merges the per-node metric streams into one stream and
// returns every sample it received, sorted ascending.
//
// Examples:
//
//	MergeMetrics(chan of 1, chan of 2)      => []int{1, 2}
//	MergeMetrics(chan of 3, 1, chan of 2)  => []int{1, 2, 3}
//	MergeMetrics()                         => nil
func MergeMetrics(streams ...<-chan int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
