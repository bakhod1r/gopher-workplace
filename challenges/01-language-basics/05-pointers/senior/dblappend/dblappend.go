// Package dblappend appends through a *[]int. A planted bug appends to a local
// dereference and discards it, so the caller's slice never grows.
package dblappend

// Extend appends all of vs to the slice pointed to by sp.
func Extend(sp *[]int, vs ...int) {
	// CHANGE CODE BELOW THIS LINE
	s := *sp
	s = append(s, vs...)
	_ = s
	// CHANGE CODE ABOVE THIS LINE
}
