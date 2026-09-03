// Package errcollector — Gopher Workplace challenge.
package errcollector

import "errors"

// Collector accumulates at most Limit errors and counts the rest.
type Collector struct {
	Limit int

	stored []error
	count  int
}

// Add records err, ignoring nil.
func (c *Collector) Add(err error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Count returns how many non-nil errors were added.
func (c *Collector) Count() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Err returns the stored errors joined, or nil when none were added.
func (c *Collector) Err() error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
