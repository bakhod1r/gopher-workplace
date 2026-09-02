// Package topkstream — Gopher Workplace challenge.
package topkstream

// Source yields values until drained.
type Source interface {
	Next() (int, bool)
}

// TopKAgg folds values, keeping the largest ones.
type TopKAgg interface {
	Add(v int)
	Result() []int
}

// SliceSource streams a slice.
type SliceSource struct {
	Data []int
	pos  int
}

// Next yields the next value.
func (s *SliceSource) Next() (int, bool) {
	if s.pos >= len(s.Data) {
		return 0, false
	}
	v := s.Data[s.pos]
	s.pos++
	return v, true
}

// TopK keeps the K largest values seen, in descending order.
type TopK struct {
	K    int
	vals []int
}

// NewTopK returns an aggregator bounded to k values.
func NewTopK(k int) *TopK {
	return &TopK{K: k, vals: make([]int, 0, k)}
}

// Add folds one value into the window.
//
// Examples:
//
//	k 2, adding 1, 5, 3 => window is [5 3]
func (t *TopK) Add(v int) {
	// TODO(candidate): insert into the descending window, keeping at most K.
	panic("not implemented")
}

// Result returns the kept values in descending order.
func (t *TopK) Result() []int {
	// TODO(candidate): return the window.
	panic("not implemented")
}

// Stream folds src through agg and returns the result.
func Stream(src Source, agg TopKAgg) []int {
	// TODO(candidate): drain, then report.
	panic("not implemented")
}
