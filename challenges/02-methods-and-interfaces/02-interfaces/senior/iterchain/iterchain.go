// Package iterchain — Gopher Workplace challenge.
package iterchain

// Iter yields values until drained.
type Iter interface {
	Next() (int, bool)
}

// SliceIter iterates a slice; Reads counts how many elements were pulled.
type SliceIter struct {
	Data  []int
	Reads int
	pos   int
}

// Next yields the next element.
func (s *SliceIter) Next() (int, bool) {
	if s.pos >= len(s.Data) {
		return 0, false
	}
	v := s.Data[s.pos]
	s.pos++
	s.Reads++
	return v, true
}

// MapFunc transforms one value.
type MapFunc func(int) int

// MapIter applies a function to every element, lazily.
type MapIter struct {
	Inner Iter
	Fn    MapFunc
}

// Next yields the next transformed element.
func (m *MapIter) Next() (int, bool) {
	// TODO(candidate): pull one, transform it.
	panic("not implemented")
}

// PredFunc reports whether a value is kept.
type PredFunc func(int) bool

// FilterIter keeps only matching elements, lazily.
type FilterIter struct {
	Inner Iter
	Pred  PredFunc
}

// Next yields the next matching element.
func (f *FilterIter) Next() (int, bool) {
	// TODO(candidate): pull until a match or the source drains.
	panic("not implemented")
}

// Collect drains it into a slice.
//
// Examples:
//
//	Collect(&MapIter{Inner: src, Fn: double}) => [2 4 6]
func Collect(it Iter) []int {
	// TODO(candidate): drain into a slice.
	panic("not implemented")
}
