// Package capgrew — Gopher Workplace challenge.
package capgrew

// Appended appends v to s and reports whether the append had to grow the
// capacity.
//
// Growing means a new array was allocated and the old contents copied.
//
// Examples:
//
//	Appended(make([]int, 0, 4), 1) => a one-element slice, false
func Appended(s []int, v int) ([]int, bool) {
	panic("not implemented")
}
