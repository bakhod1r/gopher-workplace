// Package shardrouter — Gopher Workplace challenge.
package shardrouter

// ShardIDs routes every key to a shard by hashing it.
//
// Examples:
//
//	ShardIDs([]string{"a", "b"}, 4)  => [1 2]
//	ShardIDs([]string{"a", "b"}, 1)  => [0 0]
//	ShardIDs([]string{"a"}, 0)       => []
func ShardIDs(keys []string, shards int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
