// Package shardwarmup — Gopher Workplace challenge.
package shardwarmup

// WaitForShards starts n shard warm-up goroutines that each signal
// completion on a chan struct{}, receives exactly n signals, and returns
// how many shards reported in.
//
// Examples:
//
//	WaitForShards(3) => 3
//	WaitForShards(1) => 1
//	WaitForShards(0) => 0
func WaitForShards(n int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
