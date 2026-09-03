// Package appendints — Gopher Workplace challenge.
package appendints

import "strconv"

// AppendInts renders vals as decimal numbers separated by ' ' and
// appends them to dst.
//
// Passing an int to a variadic any parameter puts it on the heap. Rendering
// must go straight into dst instead.
//
// Examples:
//
//	AppendInts(nil, []int{1, 2}) => []byte("1 2")
func AppendInts(dst []byte, vals []int) []byte {
	panic("not implemented")
}
