// Package growonce — Gopher Workplace challenge.
package growonce

// GrowTo returns a slice with the same length and contents as s but capacity
// at least n. If s already has room it is returned unchanged, without
// allocating; otherwise exactly one allocation happens.
//
// Examples:
//
//	cap(GrowTo(make([]int, 2), 100)) >= 100
func GrowTo(s []int, n int) []int {
	panic("not implemented")
}
