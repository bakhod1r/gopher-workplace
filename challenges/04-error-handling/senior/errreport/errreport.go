// Package errreport — Gopher Workplace challenge.
package errreport

import (
	"errors"
	"fmt"
	"strings"
)

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Report renders err as a numbered list, one line per failure.
//
// Examples:
//
//	Report(nil)  => ""
//	Report(ErrA) => "1. a"
func Report(err error) string {
	// TODO(candidate): implement this.
	_, _ = fmt.Sprintf, strings.Join
	panic("not implemented")
}
