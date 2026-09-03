// Package boxedintcount — Gopher Workplace challenge.
package boxedintcount

// Box converts each value to any. Storing an int in an interface normally
// copies it to the heap, but the runtime keeps a table of the small values
// 0..255, so boxing those costs nothing beyond the result slice itself.
//
// The result must be built in a single allocation.
//
// Examples:
//
//	Box([]int{1, 2}) => []any{1, 2}
func Box(xs []int) []any {
	panic("not implemented")
}

// Unbox extracts the int values again, returning the count of elements that
// were not ints along with the successfully converted ones.
//
// Examples:
//
//	Unbox([]any{1, "x"}) => []int{1}, 1
func Unbox(vs []any) ([]int, int) {
	panic("not implemented")
}
