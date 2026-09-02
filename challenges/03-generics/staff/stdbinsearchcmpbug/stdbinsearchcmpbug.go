// Package stdbinsearchcmpbug — Gopher Workplace challenge.
package stdbinsearchcmpbug

import (
	"cmp"
	"slices"
)

// Entry is one leaderboard row.
type Entry struct {
	Name  string
	Score int
}

// FindScore returns the index of the entry with the given score, and true.
// The board is sorted by Score in DESCENDING order (best first).
// It returns -1 and false when no entry has that score.
//
// Examples:
//
//	FindScore([{a 9} {b 5} {c 1}], 5) => 1, true
func FindScore(board []Entry, score int) (int, bool) {
	// CHANGE CODE BELOW THIS LINE
	i, ok := slices.BinarySearchFunc(board, score, func(e Entry, target int) int {
		return cmp.Compare(e.Score, target)
	})
	if !ok {
		return -1, false
	}
	return i, true
	// CHANGE CODE ABOVE THIS LINE
}
