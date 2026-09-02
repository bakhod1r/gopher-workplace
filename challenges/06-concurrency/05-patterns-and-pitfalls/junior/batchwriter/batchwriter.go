// Package batchwriter — Gopher Workplace challenge.
package batchwriter

// BatchInserts groups streamed rows into insert batches of at most size
// rows, flushing a final short batch if any rows remain. A size of zero or
// less returns nil.
//
// Examples:
//
//	BatchInserts(chan of a, b, c, 2)  => [][]string{{"a", "b"}, {"c"}}
//	BatchInserts(chan of a, b, 2)     => [][]string{{"a", "b"}}
//	BatchInserts(chan of a, 0)        => nil
func BatchInserts(rows <-chan string, size int) [][]string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
