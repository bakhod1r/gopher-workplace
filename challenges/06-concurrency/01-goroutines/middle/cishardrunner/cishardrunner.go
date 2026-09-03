// Package cishardrunner — Gopher Workplace challenge.
package cishardrunner

// RunShards runs every CI shard in its own goroutine and reports the total
// number of tests that passed together with the indices of the shards that
// failed, sorted ascending. A failing shard contributes no passes: a crashed
// runner cannot be trusted to have counted anything.
//
// Examples:
//
//	RunShards([][]string{{"a"}, {"b"}}, run)  => 2, []
//	RunShards([][]string{{}, {"b"}}, run)     => 1, [0]
//	RunShards(nil, run)                       => 0, []
func RunShards(shards [][]string, run func(shard []string) (passed int, err error)) (int, []int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
