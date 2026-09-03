// Package parseerr — Gopher Workplace challenge.
package parseerr

import (
	"errors"
	"fmt"
)

// ParseError reports a failure at a specific line.
type ParseError struct {
	Line int
	Msg  string
}

// Error implements the error interface as "line <Line>: <Msg>".
func (e *ParseError) Error() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf
	panic("not implemented")
}

// LineOf returns the line reported by a *ParseError in err's chain.
//
// Examples:
//
//	LineOf(&ParseError{Line: 4}) => 4, true
//	LineOf(errors.New("boom"))   => 0, false
func LineOf(err error) (int, bool) {
	// TODO(candidate): implement this.
	_ = errors.As
	panic("not implemented")
}
