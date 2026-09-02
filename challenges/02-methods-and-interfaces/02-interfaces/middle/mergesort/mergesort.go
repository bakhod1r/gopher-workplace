// Package mergesort — Gopher Workplace challenge.
package mergesort

// Feed yields ascending ints.
type Feed interface {
	// Peek returns the next value without consuming it.
	Peek() (int, bool)
	// Next consumes and returns the next value.
	Next() (int, bool)
}

// SliceFeed serves a sorted slice.
type SliceFeed struct {
	Data []int
	pos  int
}

// Peek returns the next value without consuming it.
func (f *SliceFeed) Peek() (int, bool) {
	// TODO(candidate): look at Data[pos] without advancing.
	panic("not implemented")
}

// Next consumes and returns the next value.
func (f *SliceFeed) Next() (int, bool) {
	// TODO(candidate): return Data[pos] and advance.
	panic("not implemented")
}

// Merge merges two ascending feeds into one ascending slice, in one pass.
//
// Examples:
//
//	Merge([1 3], [2 4]) => [1 2 3 4]
func Merge(a, b Feed) []int {
	// TODO(candidate): peek both, take the smaller, drain the rest.
	panic("not implemented")
}
