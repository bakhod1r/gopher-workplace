// Package semaphore — Gopher Workplace challenge.
package semaphore

// Sem is a counting semaphore of fixed capacity.
type Sem struct {
	slots chan struct{}
}

// NewSem returns a semaphore permitting n concurrent holders.
func NewSem(n int) *Sem {
	if n < 1 {
		n = 1
	}
	return &Sem{slots: make(chan struct{}, n)}
}

// Release returns one slot. It must only be called after a successful Acquire.
func (s *Sem) Release() {
	select {
	case <-s.slots:
	default:
	}
}

// Held reports how many slots are currently taken.
func (s *Sem) Held() int { return len(s.slots) }

// Acquire takes one slot, blocking until one is free, and reports whether
// it got one.
//
// It must also give up when done is closed, so a cancelled caller does not
// wait forever on a saturated semaphore.
//
// Examples:
//
//	s := NewSem(2); s.Acquire(done) => true twice, then blocks
func (s *Sem) Acquire(done <-chan struct{}) bool {
	panic("not implemented")
}
