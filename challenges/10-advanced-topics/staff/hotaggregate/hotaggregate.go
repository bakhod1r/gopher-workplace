// Package hotaggregate — Gopher Workplace challenge.
package hotaggregate

import "errors"

// ErrSyntax marks a field that is not a decimal integer.
var ErrSyntax = errors.New("invalid integer")

// Aggregate sums the decimal integers across every line and reports the
// total and the field count.
//
// The whole aggregation must run without allocating: no conversions, no
// split slices, no formatted errors.
//
// Examples:
//
//	Aggregate([][]byte{[]byte("1,2")}, ',') => 3, 2, nil
func Aggregate(lines [][]byte, sep byte) (int64, int, error) {
	panic("not implemented")
}
