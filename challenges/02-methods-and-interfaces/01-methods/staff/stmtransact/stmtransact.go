// Package stmtransact — Gopher Workplace challenge.
package stmtransact

import "sync"

// TVar is a transactional variable: a value plus the version that changes on
// every committed write.
type TVar struct {
	mu      sync.Mutex
	val     int
	version int
}

// NewTVar returns a TVar holding val at version 0.
func NewTVar(val int) *TVar {
	return &TVar{val: val}
}

// Read returns the current value and the version it was read at.
func (tv *TVar) Read() (int, int) {
	tv.mu.Lock()
	defer tv.mu.Unlock()
	return tv.val, tv.version
}

// Commit writes val only if the variable has not changed since readVersion.
// It reports whether the write happened.
func (tv *TVar) Commit(readVersion, val int) bool {
	tv.mu.Lock()
	defer tv.mu.Unlock()
	if tv.version != readVersion {
		return false
	}
	tv.val = val
	tv.version++
	return true
}

// Tx runs fn against the current value and commits the result, retrying from a
// fresh read whenever another transaction committed first. It returns the value
// that was committed.
func (tv *TVar) Tx(fn func(int) int) int {
	// TODO(candidate): loop — Read, compute fn, Commit; on a failed commit,
	// start over from a fresh Read. Return the committed value.
	panic("not implemented")
}
