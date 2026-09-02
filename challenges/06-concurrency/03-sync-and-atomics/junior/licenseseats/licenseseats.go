// Package licenseseats - Gopher Workplace challenge.
package licenseseats

import "sync/atomic"

// SeatPool hands out a bounded number of concurrent license seats.
type SeatPool struct {
	free atomic.Int64
}

// NewSeatPool returns a pool with n free seats.
func NewSeatPool(n int64) *SeatPool {
	p := &SeatPool{}
	p.free.Store(n)
	return p
}

// TryAcquire takes a seat and reports whether one was available.
//
// Examples:
//
//	p := NewSeatPool(2); p.TryAcquire()                  => true
//	p := NewSeatPool(1); p.TryAcquire(); p.TryAcquire()  => false
func (p *SeatPool) TryAcquire() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Release returns a seat to the pool.
//
// Examples:
//
//	p := NewSeatPool(1); p.TryAcquire(); p.Release(); p.Free() => 1
func (p *SeatPool) Release() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Free reports how many seats remain.
func (p *SeatPool) Free() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
