// Package parseints — Gopher Workplace challenge.
package parseints

import "errors"

// ErrSyntax is returned for a field that is not a decimal integer.
var ErrSyntax = errors.New("invalid integer")

// ParseInts sums the decimal integers in line, which are separated by
// sep, and returns the total, the count parsed, and any error.
//
// No part of line may be converted to a string: the parse works on the
// bytes and allocates nothing.
//
// Examples:
//
//	ParseInts([]byte("1,2,3"), ',') => 6, 3, nil
func ParseInts(line []byte, sep byte) (int64, int, error) {
	panic("not implemented")
}
