// Package benchbytesop — Gopher Workplace challenge.
package benchbytesop

// ThroughputMBs converts a benchmark total into the MB/s column that
// b.SetBytes produces: totalBytes moved in elapsedNS nanoseconds, using
// 1 MB = 1_000_000 bytes and 1 s = 1e9 ns. A non-positive elapsedNS is 0.
//
// Examples:
//
//	ThroughputMBs(1_000_000, 1_000_000_000) => 1
func ThroughputMBs(totalBytes int64, elapsedNS int64) float64 {
	panic("not implemented")
}
