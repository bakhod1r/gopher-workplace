// Package spinlock — Gopher Workplace challenge.
package spinlock

import (
	"runtime"
	"sync/atomic"
)

// Locker is the lock contract.
type Locker interface {
	Lock()
	Unlock()
}

// SpinLock is a spinning mutual-exclusion lock for very short sections.
type SpinLock struct {
	held atomic.Bool
}

// spinsBeforeYield bounds how long a waiter spins before yielding its P.
const spinsBeforeYield = 30

// Lock acquires the lock, spinning with a bounded backoff.
//
// Examples:
//
//	Lock(); TryLock() => false
func (l *SpinLock) Lock() {
	// TODO(candidate): CAS loop with a bounded spin, then yield.
	panic("not implemented")
}

// TryLock acquires the lock without spinning.
func (l *SpinLock) TryLock() bool {
	// TODO(candidate): a single CAS.
	panic("not implemented")
}

// Unlock releases the lock. Unlocking an unlocked SpinLock panics.
func (l *SpinLock) Unlock() {
	// TODO(candidate): release, detecting misuse.
	panic("not implemented")
}

// Locked reports whether the lock is currently held.
func (l *SpinLock) Locked() bool {
	// TODO(candidate): report the state.
	panic("not implemented")
}

var _ = runtime.Gosched
