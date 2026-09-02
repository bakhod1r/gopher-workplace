// Package cachewarm — Gopher Workplace challenge.
package cachewarm

// WarmShards starts one warmer goroutine per shard, holds them all until the
// ready channel is closed, and returns the per-shard key counts sorted
// ascending.
//
// Examples:
//
//	WarmShards(closed, []string{"a", "bb"}, keyCount)  => []int{1, 2}
//	WarmShards(closed, []string{"xyz"}, keyCount)      => []int{3}
//	WarmShards(closed, nil, keyCount)                  => []int{}
func WarmShards(ready <-chan struct{}, shards []string, warm func(string) int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
