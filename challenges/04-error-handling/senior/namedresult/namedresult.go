// Package namedresult — Gopher Workplace challenge.
package namedresult

import (
	"errors"
	"fmt"
)

// ErrBoom is a stand-in failure used by the tests.
var ErrBoom = errors.New("boom")

// Do runs f and annotates any failure with op.
//
// Examples:
//
//	Do("load", func() (int, error) { return 7, nil }) => 7, nil
func Do(op string, f func() (int, error)) (v int, err error) {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
