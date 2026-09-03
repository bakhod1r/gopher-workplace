// Package bytesopwrongbug — Gopher Workplace challenge.
package bytesopwrongbug

// ThroughputMBs renders the MB/s column b.SetBytes produces: totalBytes moved
// in elapsedNS nanoseconds, using the same units the benchmark tool does —
// 1 MB is 1,000,000 bytes and 1 s is 1,000,000,000 ns.
//
// Examples:
//
//	ThroughputMBs(1_000_000, 1_000_000_000) => 1
func ThroughputMBs(totalBytes, elapsedNS int64) float64 {
	if elapsedNS <= 0 {
		return 0
	}
	// CHANGE CODE BELOW THIS LINE
	return float64(totalBytes) / float64(elapsedNS) * 1000 * (1000000.0 / 1048576.0)
	// CHANGE CODE ABOVE THIS LINE
}

// PerOpBytes returns the bytes moved per iteration, which is what a caller
// passes to b.SetBytes.
//
// Examples:
//
//	PerOpBytes(2048, 4) => 512
func PerOpBytes(totalBytes int64, iters int64) int64 {
	if iters <= 0 {
		return 0
	}
	return totalBytes / iters
}
