// Package tolerantparse — Gopher Workplace challenge.
package tolerantparse

import (
	"errors"
	"fmt"
	"strconv"
)

// Parse converts every line, collecting per-line failures.
//
// Examples:
//
//	Parse([]string{"1"}) => [1], nil
func Parse(lines []string) ([]int, error) {
	// TODO(candidate): implement this.
	_, _, _ = fmt.Errorf, strconv.Atoi, errors.Join
	panic("not implemented")
}
