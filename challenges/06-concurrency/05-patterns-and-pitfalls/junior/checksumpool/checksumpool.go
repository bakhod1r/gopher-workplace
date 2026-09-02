// Package checksumpool — Gopher Workplace challenge.
package checksumpool

// result pairs a file with the checksum a worker computed for it.
type result struct {
	file string
	sum  int
}

// ChecksumFiles checksums every file through a fixed worker pool. Each worker
// reports a file/sum pair on the results channel, and the caller builds the
// map alone. workers is >= 1.
//
// Examples:
//
//	ChecksumFiles([]string{"a"}, 2, sum)       => map[a:1]
//	ChecksumFiles([]string{"a", "bb"}, 2, sum)  => map[a:1 bb:2]
//	ChecksumFiles(nil, 3, sum)                  => empty map
func ChecksumFiles(files []string, workers int, sum func(string) int) map[string]int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
