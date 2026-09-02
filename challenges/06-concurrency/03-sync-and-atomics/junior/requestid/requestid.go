// Package requestid - Gopher Workplace challenge.
package requestid

import "sync/atomic"

// IDGen hands out unique, increasing request IDs.
type IDGen struct {
	n atomic.Int64
}

// Next returns the next unique request ID, starting at 1.
//
// Examples:
//
//	var g IDGen; g.Next()           => 1
//	var g IDGen; g.Next(); g.Next() => 2
func (g *IDGen) Next() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Issued reports how many IDs have been handed out.
//
// Examples:
//
//	var g IDGen; g.Next(); g.Issued() => 1
func (g *IDGen) Issued() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
