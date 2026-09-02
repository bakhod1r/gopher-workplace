// Package cyclicbarrier — Gopher Workplace challenge.
package cyclicbarrier

import "sync"

// Waiter is a rendezvous point.
type Waiter interface {
	Await() int
}

// Barrier releases goroutines in groups of Parties.
type Barrier struct {
	Parties int

	mu    sync.Mutex
	cond  *sync.Cond
	count int
	round int
}

// NewBarrier returns a barrier for the given number of parties.
func NewBarrier(parties int) *Barrier {
	if parties < 1 {
		parties = 1
	}
	b := &Barrier{Parties: parties}
	b.cond = sync.NewCond(&b.mu)
	return b
}

// Await blocks until Parties goroutines have arrived, then returns the index
// of the round that just completed.
//
// Examples:
//
//	3 parties; three Await calls => all three return 0
func (b *Barrier) Await() int {
	// TODO(candidate): arrive; the last one rearms and broadcasts.
	panic("not implemented")
}

// Round returns how many rounds have completed.
func (b *Barrier) Round() int {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}
