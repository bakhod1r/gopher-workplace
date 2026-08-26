// Package iterator — Gopher Workplace challenge.
package iterator

// IntIter iterates over a slice of ints.
type IntIter struct {
	data []int
	pos  int
}

// NewIntIter creates an iterator over the given slice.
func NewIntIter(data []int) *IntIter {
	return &IntIter{data: data, pos: 0}
}

// Next advances the iterator and returns true if there is a value to read.
//
// Examples:
//
//	it := NewIntIter([]int{10, 20})
//	it.Next() => true;  it.Value() => 10
//	it.Next() => true;  it.Value() => 20
//	it.Next() => false
func (it *IntIter) Next() bool {
	// TODO(candidate): advance position, return whether valid.
	panic("not implemented")
}

// Value returns the current element. Call only after Next returns true.
func (it *IntIter) Value() int {
	// TODO(candidate): return the current element.
	panic("not implemented")
}
