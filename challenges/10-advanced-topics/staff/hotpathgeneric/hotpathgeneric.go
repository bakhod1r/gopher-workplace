// Package hotpathgeneric — Gopher Workplace challenge.
package hotpathgeneric

// Total sums vals.
//
// An `[]any` version would box every element. A type parameter gives the
// compiler the concrete type, so nothing is boxed and nothing escapes.
//
// Examples:
//
//	Total([]int{1, 2, 3}) => 6
func Total[T ~int | ~int32 | ~int64](vals []T) int64 {
	panic("not implemented")
}
