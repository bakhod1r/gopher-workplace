// Package fielderrs — Gopher Workplace challenge.
package fielderrs

import (
	"errors"
	"fmt"
	"sort"
)

// Field failures used by the tests.
var (
	ErrA = errors.New("bad a")
	ErrB = errors.New("bad b")
)

// Combine merges per-field failures into one error, fields in sorted order.
//
// Examples:
//
//	Combine(map[string]error{"a": ErrA}) => "a: bad a"
//	Combine(nil)                         => nil
func Combine(m map[string]error) error {
	// TODO(candidate): implement this.
	_, _, _ = fmt.Errorf, sort.Strings, errors.Join
	panic("not implemented")
}
