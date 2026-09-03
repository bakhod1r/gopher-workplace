// Package indexwrap — Gopher Workplace challenge.
package indexwrap

import (
	"errors"
	"fmt"
)

// ErrParse is a stand-in parse failure used by the tests.
var ErrParse = errors.New("parse failed")

// AtIndex annotates err with the record index.
//
// Examples:
//
//	AtIndex(3, ErrParse) => "record 3: parse failed"
//	AtIndex(0, nil)      => nil
func AtIndex(i int, err error) error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
