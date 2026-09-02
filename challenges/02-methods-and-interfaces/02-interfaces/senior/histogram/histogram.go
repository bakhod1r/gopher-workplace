// Package histogram — Gopher Workplace challenge.
package histogram

// Recorder observes samples.
type Recorder interface {
	Observe(v int)
}

// Histogram counts samples into fixed buckets.
//
// Bounds must be ascending. A value v lands in the first bucket whose bound
// is >= v; values above every bound land in the overflow bucket.
type Histogram struct {
	Bounds   []int
	counts   []int
	overflow int
	total    int
}

// NewHistogram returns a histogram over the given ascending bounds.
func NewHistogram(bounds []int) *Histogram {
	return &Histogram{Bounds: bounds, counts: make([]int, len(bounds))}
}

// Observe records one sample.
//
// Examples:
//
//	bounds [10 100]; Observe(5) => bucket 0
func (h *Histogram) Observe(v int) {
	// TODO(candidate): find the bucket, count it, count the total.
	panic("not implemented")
}

// Count returns how many samples were observed.
func (h *Histogram) Count() int {
	// TODO(candidate): total samples.
	panic("not implemented")
}

// Quantile returns the upper bound of the bucket holding the q-th sample
// (0 <= q <= 1). It returns 0 for an empty histogram, and the last bound
// when the quantile falls in the overflow bucket.
func (h *Histogram) Quantile(q float64) int {
	// TODO(candidate): prefix scan over the bucket counts.
	panic("not implemented")
}
