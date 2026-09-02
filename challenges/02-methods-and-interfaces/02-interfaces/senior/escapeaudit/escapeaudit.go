// Package escapeaudit — Gopher Workplace challenge.
package escapeaudit

// Summer sums a slice of ints.
type Summer interface {
	Sum(vs []int) int
}

// StackAgg sums without letting its state escape.
type StackAgg struct{}

// Sum adds up the values.
//
// Examples:
//
//	StackAgg{}.Sum([]int{1, 2, 3}) => 6
func (StackAgg) Sum(vs []int) int {
	// TODO(candidate): a plain loop over the concrete slice.
	panic("not implemented")
}

// SumValues sums a concrete slice with no boxing.
func SumValues(vs []int) int {
	// TODO(candidate): no interface conversions.
	panic("not implemented")
}

// SumBoxed sums values that have already been boxed into any.
// Elements that are not ints are skipped.
func SumBoxed(vs []any) int {
	// TODO(candidate): comma-ok assertion per element.
	panic("not implemented")
}

// BoxAll converts a concrete slice into boxed values.
// Every element boxed here escapes to the heap.
func BoxAll(vs []int) []any {
	// TODO(candidate): one boxed element per input.
	panic("not implemented")
}
