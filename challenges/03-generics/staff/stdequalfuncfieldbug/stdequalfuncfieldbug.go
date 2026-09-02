// Package stdequalfuncfieldbug — Gopher Workplace challenge.
package stdequalfuncfieldbug

import (
	"slices"
)

// Line is one line of an order. Note is free text and carries no identity.
type Line struct {
	SKU  string
	Qty  int
	Note string
}

// SameLines reports whether two order line-ups are equivalent.
// Two lines are equivalent when SKU and Qty match; Note is free text and is ignored.
//
// Examples:
//
//	SameLines([{x 1 "a"}], [{x 1 "b"}]) => true
func SameLines(a, b []Line) bool {
	// CHANGE CODE BELOW THIS LINE
	return slices.Equal(a, b)
	// CHANGE CODE ABOVE THIS LINE
}
