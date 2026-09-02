// Package shardsplit — Gopher Workplace challenge.
package shardsplit

// SplitByShard routes user ids onto two delivery queues by parity — even
// ids to shard 0, odd ids to shard 1 — closes both, and returns the drained
// queues in input order.
//
// Each result slice is non-nil.
//
// Examples:
//
//	SplitByShard([]int{1,2,3,4}) => [2 4], [1 3]
//	SplitByShard([]int{2}) => [2], []
//	SplitByShard(nil) => [], []
func SplitByShard(userIDs []int) ([]int, []int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
