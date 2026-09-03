// Package wraponce — Gopher Workplace challenge.
package wraponce

import (
	"errors"
	"fmt"
	"strings"
)

// ErrA is a stand-in failure used by the tests.
var ErrA = errors.New("a")

// Once annotates err with op unless the prefix is already present.
//
// Examples:
//
//	Once("save", nil) => nil
func Once(op string, err error) error {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, strings.HasPrefix
	panic("not implemented")
}
