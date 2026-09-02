// Package shardids — Gopher Workplace challenge.
package shardids

// StreamShardIDs returns a channel that yields shard ids 0, 1, ... n-1
// and is then closed. A non-positive n produces an already-closed
// empty channel.
//
// Examples:
//
//	StreamShardIDs(3)   => 0, 1, 2 then closed
//	StreamShardIDs(1)   => 0 then closed
//	StreamShardIDs(0)   => closed immediately
func StreamShardIDs(n int) <-chan int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
