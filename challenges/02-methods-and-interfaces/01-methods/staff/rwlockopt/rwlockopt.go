// Package rwlockopt — Gopher Workplace challenge.
package rwlockopt

import "sync"

// OptLock prefers read locks but escalates to write locks if needed.
type OptLock struct {
	mu sync.RWMutex
	v  int
}

// IncrementIfZero reads the value. If zero, it escalates to a write lock
// to increment it. Returns the final value.
func (o *OptLock) IncrementIfZero() int {
	// TODO(candidate):
	// RLock, read v. If v != 0, RUnlock, return v.
	// RUnlock. Lock. If v == 0 { v++ }. Unlock. return v.
	panic("not implemented")
}
