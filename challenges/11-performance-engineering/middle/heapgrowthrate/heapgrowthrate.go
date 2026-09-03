// Package heapgrowthrate — Gopher Workplace challenge.
package heapgrowthrate

// Sample is one heap observation: live bytes at a moment in time.
type Sample struct {
	NS   int64
	Live int64
}

// NextTarget returns the heap size at which the next collection starts, given
// the live heap after the last one and the GOGC percentage: GOGC=100 means
// collect when the heap has doubled. A GOGC of 0 or less disables the
// percentage trigger and reports false.
//
// Examples:
//
//	NextTarget(4<<20, 100) => 8<<20, true
func NextTarget(liveBytes int64, gogc int) (int64, bool) {
	panic("not implemented")
}

// GrowthPerSec returns the bytes-per-second slope between the first and last
// samples — a steadily positive value across collections is the signature of
// a leak rather than of ordinary allocation. Fewer than two samples, or a
// non-increasing timestamp, gives 0 and false.
//
// Examples:
//
//	GrowthPerSec([{0, 100}, {1e9, 300}]) => 200, true
func GrowthPerSec(samples []Sample) (float64, bool) {
	panic("not implemented")
}

// Doubling returns how many seconds the heap takes to double at that slope,
// starting from the last sample. A non-positive slope never doubles and
// reports false.
//
// Examples:
//
//	Doubling([{0, 100}, {1e9, 200}]) => 2, true
func Doubling(samples []Sample) (float64, bool) {
	panic("not implemented")
}
