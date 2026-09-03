// Package temperr — Gopher Workplace challenge.
package temperr

import "errors"

// NetError is a transport failure that knows whether it is retryable.
type NetError struct {
	Temp bool
}

// Error implements the error interface.
func (e *NetError) Error() string {
	return "network failure"
}

// Temporary reports whether the failure is retryable.
func (e *NetError) Temporary() bool {
	return e.Temp
}

// IsTemporary reports whether err advertises itself as temporary.
//
// Examples:
//
//	IsTemporary(&NetError{Temp: true})  => true
//	IsTemporary(&NetError{Temp: false}) => false
func IsTemporary(err error) bool {
	// TODO(candidate): implement this.
	_ = errors.As
	panic("not implemented")
}
