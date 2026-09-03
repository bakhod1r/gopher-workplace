// Package safeappend — Gopher Workplace challenge.
package safeappend

// Add returns s with v appended.
//
// The caller may be holding a longer slice over the same array. Appending
// must never overwrite elements past len(s); the result must get its own
// storage whenever that would happen.
//
// Examples:
//
//	Add([]int{1, 2}, 3) => []int{1, 2, 3}
func Add(s []int, v int) []int {
	// CHANGE CODE BELOW THIS LINE
	return append(s, v)
	// CHANGE CODE ABOVE THIS LINE
}
