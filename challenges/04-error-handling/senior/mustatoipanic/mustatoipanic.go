// Package mustatoipanic — Gopher Workplace challenge.
package mustatoipanic

import (
	"fmt"
	"strconv"
)

// MustParse converts s to an int, panicking on failure.
//
// Examples:
//
//	MustParse("42") => 42
//	MustParse("x")  => panics
func MustParse(s string) int {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, strconv.Atoi
	panic("not implemented")
}
