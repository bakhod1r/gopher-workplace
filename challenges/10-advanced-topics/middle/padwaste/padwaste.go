// Package padwaste — Gopher Workplace challenge.
package padwaste

import "reflect"

// Waste returns how many bytes of v's struct type are padding: its total
// size minus the sum of its fields' sizes.
//
// A non-struct wastes nothing.
//
// Examples:
//
//	Waste(gappy{}) => 14 for a byte, an int64 and a byte
func Waste(v any) uintptr {
	panic("not implemented")
}
