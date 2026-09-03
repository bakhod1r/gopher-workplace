// Package warehouseslots — Gopher Workplace challenge.
package warehouseslots

import "sync/atomic"

// Stock is the reservable unit count of one SKU in a fulfilment centre. Many
// checkout goroutines race to reserve the last units; overselling is not
// recoverable, so the count must never drop below zero.
type Stock struct {
	available atomic.Int64
}

// NewStock returns a Stock holding units; a negative count is treated as 0.
//
// Examples:
//
//	NewStock(10).Available() => 10
//	NewStock(-3).Available()  => 0
func NewStock(units int64) *Stock {
	s := &Stock{}
	if units > 0 {
		s.available.Store(units)
	}
	return s
}

// Reserve takes n units and reports whether the reservation succeeded. It
// fails when n is not positive or when fewer than n units remain. It is a
// compare-and-swap retry loop: read, decide, swap, and start over if another
// goroutine moved the count in between.
//
// Examples:
//
//	s := NewStock(10); s.Reserve(3)  => true
//	s := NewStock(2); s.Reserve(5)   => false
//	NewStock(10).Reserve(0)          => false
func (s *Stock) Reserve(n int64) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Release returns n units to the pool. A non-positive n is ignored.
//
// Examples:
//
//	s := NewStock(1); s.Reserve(1); s.Release(1); s.Available() => 1
func (s *Stock) Release(n int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Available returns the units currently reservable.
func (s *Stock) Available() int64 {
	return s.available.Load()
}
