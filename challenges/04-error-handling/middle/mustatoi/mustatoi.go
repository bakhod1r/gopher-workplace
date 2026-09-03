// Package mustatoi — Gopher Workplace challenge.
package mustatoi

import (
	"fmt"
	"strconv"
)

// ParseInt converts s to an int, annotating failures with the input.
//
// Examples:
//
//	ParseInt("42") => 42, nil
//	ParseInt("x")  => 0, `parse "x": ...`
func ParseInt(s string) (int, error) {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, strconv.Atoi
	panic("not implemented")
}
