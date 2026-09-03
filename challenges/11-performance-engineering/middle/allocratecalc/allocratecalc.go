// Package allocratecalc — Gopher Workplace challenge.
package allocratecalc

// Stats is a snapshot of the cumulative allocation counters runtime.MemStats
// exposes: TotalAlloc and Mallocs only ever increase.
type Stats struct {
	NS         int64
	TotalAlloc uint64
	Mallocs    uint64
	Frees      uint64
}

// Delta returns the bytes allocated, objects allocated and objects freed
// between two snapshots. Because the counters are cumulative and monotonic, a
// later snapshot with a smaller counter means the process restarted — the
// delta is meaningless and reports false, as does a non-increasing timestamp.
//
// Examples:
//
//	Delta(a, b) => bytes, mallocs, frees, true
func Delta(from, to Stats) (bytes, mallocs, frees uint64, ok bool) {
	panic("not implemented")
}

// BytesPerSec returns the allocation rate between two snapshots, in bytes per
// second, and false when the delta is not usable.
//
// Examples:
//
//	BytesPerSec(a, b) => 1_000_000, true
func BytesPerSec(from, to Stats) (float64, bool) {
	panic("not implemented")
}

// LiveObjects returns how many objects were allocated and not yet freed at a
// snapshot: Mallocs minus Frees. A snapshot with more frees than mallocs is
// impossible and reports false.
//
// Examples:
//
//	LiveObjects(Stats{Mallocs: 10, Frees: 4}) => 6, true
func LiveObjects(s Stats) (uint64, bool) {
	panic("not implemented")
}
