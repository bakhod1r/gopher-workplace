// Package stdcompactadjbug — Gopher Workplace challenge.
package stdcompactadjbug

import (
	"slices"
	"strings"
)

// DistinctFold returns the distinct values of s, comparing case-insensitively.
// The result is ordered case-insensitively and keeps the first spelling seen.
//
// Examples:
//
//	DistinctFold([]string{"b", "B", "a"}) => []string{"a", "b"}
func DistinctFold(s []string) []string {
	// CHANGE CODE BELOW THIS LINE
	out := slices.Clone(s)
	return slices.CompactFunc(out, strings.EqualFold)
	// CHANGE CODE ABOVE THIS LINE
}
