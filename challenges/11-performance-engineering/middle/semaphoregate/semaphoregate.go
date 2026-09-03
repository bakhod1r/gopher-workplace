// Package semaphoregate — Gopher Workplace challenge.
package semaphoregate

// Sema is a counting semaphore built on a buffered channel: the buffer size
// is the number of permits, and holding a slot is holding a permit.
type Sema struct {
	slots chan struct{}
}

// New returns a semaphore with n permits. A non-positive n gives one permit.
//
// Examples:
//
//	New(4)
func New(n int) *Sema {
	panic("not implemented")
}

// Acquire takes a permit, blocking until one is free.
//
// Examples:
//
//	s.Acquire(); defer s.Release()
func (s *Sema) Acquire() {
	panic("not implemented")
}

// TryAcquire takes a permit if one is free right now and reports whether it
// did. It never blocks.
//
// Examples:
//
//	if s.TryAcquire() { defer s.Release() }
func (s *Sema) TryAcquire() bool {
	panic("not implemented")
}

// Release returns a permit. Releasing a permit that was never acquired is a
// programming error, and this implementation must not block or corrupt the
// count when it happens: the extra release is dropped.
//
// Examples:
//
//	s.Release()
func (s *Sema) Release() {
	panic("not implemented")
}

// Available reports how many permits are free right now.
//
// Examples:
//
//	New(4).Available() => 4
func (s *Sema) Available() int {
	panic("not implemented")
}
