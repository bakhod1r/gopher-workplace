// Package interncache — Gopher Workplace challenge.
package interncache

import "unsafe"

// Pool holds one canonical string per distinct byte sequence.
type Pool struct {
	m map[string]string
}

// Len reports how many distinct strings the pool holds.
func (p *Pool) Len() int { return len(p.m) }

// Intern returns a string with b's contents, reusing a previously stored
// one when the same bytes have been seen before.
//
// Repeated values then share one allocation instead of one each, and a
// repeat lookup must not allocate at all.
//
// Examples:
//
//	p.Intern([]byte("a")) twice => the same string, one allocation
func (p *Pool) Intern(b []byte) string {
	panic("not implemented")
}
