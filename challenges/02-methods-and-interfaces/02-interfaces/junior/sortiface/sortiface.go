// Package sortiface — Gopher Workplace challenge.
package sortiface

// Sortable is everything a sort algorithm needs.
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// IntSlice is a sortable list of ints.
type IntSlice []int

// Len returns the number of elements.
func (s IntSlice) Len() int {
	// TODO(candidate): length.
	panic("not implemented")
}

// Less reports whether element i sorts before element j.
//
// Examples:
//
//	IntSlice{3, 1}.Less(1, 0) => true
func (s IntSlice) Less(i, j int) bool {
	// TODO(candidate): ascending order.
	panic("not implemented")
}

// Swap exchanges elements i and j.
func (s IntSlice) Swap(i, j int) {
	// TODO(candidate): exchange in place.
	panic("not implemented")
}

// BubbleSort sorts data in place using only the interface methods.
func BubbleSort(data Sortable) {
	// TODO(candidate): repeatedly swap out-of-order neighbours.
	panic("not implemented")
}
