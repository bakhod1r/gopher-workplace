// Package growslice — Gopher Workplace challenge.
package growslice

// Grow returns s with capacity for at least n more elements, without
// changing its length or contents.
//
// If s already has the room, it is returned untouched and nothing is
// allocated. n < 0 is treated as 0.
//
// Examples:
//
//	Grow(make([]int, 2, 2), 8) => length 2, capacity at least 10
func Grow(s []int, n int) []int {
	panic("not implemented")
}
