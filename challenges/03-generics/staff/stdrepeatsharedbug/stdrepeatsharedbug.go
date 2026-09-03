// Package stdrepeatsharedbug — Gopher Workplace challenge.
package stdrepeatsharedbug

import (
	"slices"
)

// Blank returns n rows, each a copy of proto.
// Every row is independent: writing to one must not affect the others,
// and none of them may alias proto. It returns an empty result for n <= 0.
//
// Examples:
//
//	Blank([]int{0, 0}, 3) => three independent [0 0] rows
func Blank[T any](proto []T, n int) [][]T {
	// CHANGE CODE BELOW THIS LINE
	if n <= 0 {
		return [][]T{}
	}
	return slices.Repeat([][]T{proto}, n)
	// CHANGE CODE ABOVE THIS LINE
}
