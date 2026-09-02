// Package backupupload — Gopher Workplace challenge.
package backupupload

// TotalUploaded uploads every backup shard through a fixed worker pool and
// returns the sum of the bytes reported by each upload. workers is >= 1.
//
// Examples:
//
//	TotalUploaded([]string{"a", "bb"}, 2, size)  => 3
//	TotalUploaded([]string{"abcd"}, 1, size)     => 4
//	TotalUploaded(nil, 3, size)                  => 0
func TotalUploaded(shards []string, workers int, upload func(string) int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
