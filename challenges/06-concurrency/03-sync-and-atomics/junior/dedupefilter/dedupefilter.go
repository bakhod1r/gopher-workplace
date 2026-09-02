// Package dedupefilter - Gopher Workplace challenge.
package dedupefilter

import "sync"

// DedupeFilter accepts each webhook event ID exactly once.
type DedupeFilter struct {
	seen sync.Map
}

// Accept reports whether this delivery is the first for eventID.
//
// Examples:
//
//	var d DedupeFilter; d.Accept("evt-1") => true
//	d.Accept("evt-1")                     => false
func (d *DedupeFilter) Accept(eventID string) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Seen reports whether eventID has ever been accepted.
//
// Examples:
//
//	var d DedupeFilter; d.Seen("evt-1") => false
func (d *DedupeFilter) Seen(eventID string) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Len returns the number of distinct accepted events.
//
// Examples:
//
//	d.Accept("evt-1"); d.Accept("evt-2"); d.Len() => 2
func (d *DedupeFilter) Len() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
