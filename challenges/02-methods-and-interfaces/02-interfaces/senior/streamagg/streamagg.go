// Package streamagg — Gopher Workplace challenge.
package streamagg

// Source yields readings until it is drained.
type Source interface {
	Next() (int, bool)
}

// Aggregator folds readings into a single result.
type Aggregator interface {
	Add(v int)
	Result() int
}

// RangeSource yields 1..N without materialising them.
type RangeSource struct {
	N   int
	pos int
}

// Next yields the next reading.
func (r *RangeSource) Next() (int, bool) {
	if r.pos >= r.N {
		return 0, false
	}
	r.pos++
	return r.pos, true
}

// MeanAgg computes a truncated mean in constant memory.
type MeanAgg struct {
	sum   int
	count int
}

// Add folds one reading in.
func (m *MeanAgg) Add(v int) {
	// TODO(candidate): update the running state only.
	panic("not implemented")
}

// Result returns the truncated mean, or 0 for an empty stream.
func (m *MeanAgg) Result() int {
	// TODO(candidate): guard the empty case.
	panic("not implemented")
}

// MaxAgg tracks the largest reading.
type MaxAgg struct {
	max    int
	seenIt bool
}

// Add folds one reading in.
func (m *MaxAgg) Add(v int) {
	// TODO(candidate): track the maximum.
	panic("not implemented")
}

// Result returns the maximum, or 0 for an empty stream.
func (m *MaxAgg) Result() int {
	// TODO(candidate): guard the empty case.
	panic("not implemented")
}

// Aggregate drains src through agg and returns the result.
//
// Examples:
//
//	Aggregate(&RangeSource{N: 3}, &MeanAgg{}) => 2
func Aggregate(src Source, agg Aggregator) int {
	// TODO(candidate): stream; never buffer the readings.
	panic("not implemented")
}
