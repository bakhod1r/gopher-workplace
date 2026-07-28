// Package copyslice makes an independent copy of a slice.
package copyslice

// Clone returns a new slice with the same elements as xs, sharing no backing
// array. Cloning nil returns an empty (non-nil) slice.
//
// TODO(candidate): use make + copy (or append).
func Clone(xs []int) []int {
	panic("not implemented")
}
