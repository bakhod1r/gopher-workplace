// Package arrayval — Gopher Workplace challenge.
package arrayval

// SetFirst returns a copy of the array a with its first element replaced by v.
// A Go array is a value: passing it to a function copies it, so the caller's
// original array must stay unchanged — the new value comes back as the result.
//
// Examples:
//
//	SetFirst([3]int{1, 2, 3}, 9) => [3]int{9, 2, 3}   // caller's array untouched
//	SetFirst([3]int{0, 0, 0}, 5) => [3]int{5, 0, 0}
func SetFirst(a [3]int, v int) [3]int {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
