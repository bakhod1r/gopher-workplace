// Package deploybarrier — Gopher Workplace challenge.
package deploybarrier

import "sync"

// Barrier releases its participants only once all of them have arrived, then
// resets itself for the next phase.
type Barrier struct {
	mu      sync.Mutex
	cond    *sync.Cond
	need    int
	arrived int
	phase   int
}

// NewBarrier returns a Barrier for n participants; n <= 0 collapses to 1.
//
// Examples:
//
//	NewBarrier(3).Phase() => 0
func NewBarrier(n int) *Barrier {
	// TODO(candidate): build the Barrier and attach a Cond to its mutex.
	panic("not implemented")
}

// Wait blocks until every participant of the current phase has called Wait.
// The last arrival releases them all and advances the phase.
//
// Examples:
//
//	b := NewBarrier(1); b.Wait(); b.Phase() => 1
func (b *Barrier) Wait() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Phase returns how many phases have completed.
func (b *Barrier) Phase() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.phase
}
