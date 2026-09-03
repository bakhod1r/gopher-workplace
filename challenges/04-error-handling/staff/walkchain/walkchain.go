// Package walkchain — Gopher Workplace challenge.
package walkchain

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Walk calls visit for every error in err's tree, pre-order.
//
// Examples:
//
//	Walk(nil, f) => no calls
func Walk(err error, visit func(error)) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
