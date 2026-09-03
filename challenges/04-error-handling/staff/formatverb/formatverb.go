// Package formatverb — Gopher Workplace challenge.
package formatverb

import (
	"fmt"
	"io"
)

// DetailError carries a short message and a verbose detail.
type DetailError struct {
	Msg    string
	Detail string
}

// Error returns the short message.
func (e *DetailError) Error() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Format renders %v and %s as the message, and %+v with the detail.
func (e *DetailError) Format(s fmt.State, verb rune) {
	// TODO(candidate): implement this.
	_ = io.WriteString
	panic("not implemented")
}
