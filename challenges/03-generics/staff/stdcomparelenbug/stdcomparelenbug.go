// Package stdcomparelenbug — Gopher Workplace challenge.
package stdcomparelenbug

import (
	"slices"
)

// ComparePaths orders two paths shortest-first.
// A shorter path always sorts before a longer one; equal lengths compare
// segment by segment. It returns -1, 0, or +1.
//
// Examples:
//
//	ComparePaths([]int{9}, []int{1, 2}) => -1
func ComparePaths(a, b []int) int {
	// CHANGE CODE BELOW THIS LINE
	return slices.Compare(a, b)
	// CHANGE CODE ABOVE THIS LINE
}
