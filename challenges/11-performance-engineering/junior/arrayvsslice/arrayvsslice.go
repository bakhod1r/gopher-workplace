// Package arrayvsslice — Gopher Workplace challenge.
package arrayvsslice

// Block is a fixed-size array type. Passing one copies all eight elements.
type Block [8]int

// SumBlock returns the sum of the block. The parameter is a copy, so writing
// to it cannot affect the caller.
//
// Examples:
//
//	SumBlock(Block{1, 2}) => 3
func SumBlock(b Block) int {
	panic("not implemented")
}

// ZeroBlock sets every element of the caller's copy to 0 and returns it,
// leaving the argument untouched.
//
// Examples:
//
//	ZeroBlock(Block{1, 2}) => Block{}
func ZeroBlock(b Block) Block {
	panic("not implemented")
}

// SumSlice returns the sum of a slice, which is passed as a header pointing
// at the caller's array.
//
// Examples:
//
//	SumSlice([]int{1, 2}) => 3
func SumSlice(s []int) int {
	panic("not implemented")
}
