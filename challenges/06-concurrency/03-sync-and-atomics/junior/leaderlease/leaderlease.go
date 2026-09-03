// Package leaderlease — Gopher Workplace challenge.
package leaderlease

import "sync/atomic"

// Lease is the leader lock of a scheduler cluster. Exactly one replica may
// win each term, and terms only move forward.
type Lease struct {
	term atomic.Int64
}

// Claim tries to move the lease from term-1 to term. It reports whether this
// caller won the term.
//
// Examples:
//
//	var l Lease; l.Claim(1)            => true
//	var l Lease; l.Claim(1); l.Claim(1) => false
//	var l Lease; l.Claim(5)            => false
func (l *Lease) Claim(term int64) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Term returns the current term.
//
// Examples:
//
//	var l Lease; l.Term() => 0
func (l *Lease) Term() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
