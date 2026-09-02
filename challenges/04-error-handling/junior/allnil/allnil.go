// Package allnil — Gopher Workplace challenge.
package allnil

import "errors"

// ErrCheck is a stand-in failure used by the tests.
var ErrCheck = errors.New("check failed")

// AllNil reports whether every entry of errs is nil.
//
// Examples:
//
//	AllNil([]error{nil, nil})     => true
//	AllNil([]error{nil, ErrCheck}) => false
func AllNil(errs []error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
