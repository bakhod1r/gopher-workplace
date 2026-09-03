// Package tdigestlite — Gopher Workplace challenge.
package tdigestlite

// Sketch estimates quantiles in fixed memory: it never stores samples, only
// per-bucket counts against ascending upper bounds.
type Sketch struct {
	Bounds []float64
	counts []int64
	total  int64
}

// New returns a sketch over the given ascending bounds. It has len(bounds)+1
// buckets, the last holding everything past the final bound.
//
// Examples:
//
//	New([]float64{1, 10, 100})
func New(bounds []float64) *Sketch {
	panic("not implemented")
}

// Add observes one value.
//
// Examples:
//
//	s.Add(5)
func (s *Sketch) Add(v float64) {
	panic("not implemented")
}

// Count returns how many values have been observed.
//
// Examples:
//
//	s.Count() => 3
func (s *Sketch) Count() int64 {
	panic("not implemented")
}

// Quantile estimates the p-th percentile as the upper bound of the first
// bucket whose cumulative count reaches p percent of the total — so the answer
// is always one of the bounds, and it is an upper estimate. p is clamped into
// [0,100]; falling in the overflow bucket, or having observed nothing, gives
// 0, false.
//
// Examples:
//
//	s.Quantile(50) => a bound, true
func (s *Sketch) Quantile(p float64) (float64, bool) {
	panic("not implemented")
}
