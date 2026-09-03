// Package fieldptr — Gopher Workplace challenge.
package fieldptr

import "unsafe"

// Rec is a record whose Seq field is reached by offset.
type Rec struct {
	Tag  byte
	Seq  int64
	Name string
}

// BumpSeq increments the Seq field of the record p points at, using the
// field's offset rather than the field selector, and returns the new value.
//
// This is what a generic marshaller does when it only knows the offset.
//
// Examples:
//
//	r := &Rec{Seq: 1}; BumpSeq(r) => 2, r.Seq is 2
func BumpSeq(p *Rec) int64 {
	panic("not implemented")
}
