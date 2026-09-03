// Package fieldpad — Gopher Workplace challenge.
package fieldpad

// AlignUp rounds n up to the next multiple of a, the rule the compiler
// applies to every field offset. An a of 0 or 1 leaves n alone.
//
// Examples:
//
//	AlignUp(9, 8) => 16
func AlignUp(n, a int) int {
	panic("not implemented")
}

// StructSize computes the size of a struct whose fields have the given sizes,
// where each field is aligned to its own size and the whole struct is padded
// to the largest field's alignment. Sizes of zero are skipped; no fields is 0.
//
// Examples:
//
//	StructSize([]int{1, 8}) => 16
func StructSize(sizes []int) int {
	panic("not implemented")
}
