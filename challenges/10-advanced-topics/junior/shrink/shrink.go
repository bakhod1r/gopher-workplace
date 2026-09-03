// Package shrink — Gopher Workplace challenge.
package shrink

// Shrink returns a copy of s sized exactly to its length when s is
// holding on to far more capacity than it uses, and returns s unchanged
// otherwise.
//
// "Far more" means the capacity is more than twice the length.
//
// Examples:
//
//	Shrink(make([]int, 2, 64)) => a slice of length 2 and capacity 2
func Shrink(s []int) []int {
	panic("not implemented")
}
