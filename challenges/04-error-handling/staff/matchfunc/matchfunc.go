// Package matchfunc — Gopher Workplace challenge.
package matchfunc

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("timeout")
)

// Any reports whether pred holds for any error in err's tree.
//
// Examples:
//
//	Any(nil, func(error) bool { return true }) => false
func Any(err error, pred func(error) bool) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
