// Package stringer — Gopher Workplace challenge.
package stringer

import "fmt"

// Color represents an RGB color.
type Color struct {
	R, G, B uint8
}

// String implements fmt.Stringer, returning "#RRGGBB".
//
// Examples:
//
//	Color{255, 0, 128}.String() => "#ff0080"
func (c Color) String() string {
	_ = fmt.Sprintf
	// TODO(candidate): return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	panic("not implemented")
}
