// Package shardbackup — Gopher Workplace challenge.
package shardbackup

// Shard is one piece of a nightly database backup.
type Shard struct {
	ID   string
	Size int
}

// UploadShards uploads every shard in its own goroutine and returns one slot
// per shard, in input order: nil when that shard landed, the upload error when
// it did not. Every shard is attempted even after an earlier one fails, so a
// single bad object store key never truncates the backup set.
//
// Examples:
//
//	UploadShards([]Shard{{"a", 1}, {"b", 0}}, upload)  => [<nil> errEmptyShard]
//	UploadShards([]Shard{{"a", 1}}, upload)            => [<nil>]
//	UploadShards(nil, upload)                          => []
func UploadShards(shards []Shard, upload func(Shard) error) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
