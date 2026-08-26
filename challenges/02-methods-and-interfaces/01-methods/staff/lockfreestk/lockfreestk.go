// Package lockfreestk — Gopher Workplace challenge.
package lockfreestk

import "sync/atomic"

type node struct {
	val  int
	next *node
}

// Stack is a lock-free stack using atomic.Pointer (or atomic.Value for older Go).
// Here we just use a simplified spin-lock to simulate lock-free logic.
type Stack struct {
	head atomic.Pointer[node]
}

// Push pushes a value atomically.
func (s *Stack) Push(val int) {
	n := &node{val: val}
	_ = n // hint
	// TODO(candidate): spin with CompareAndSwap until successful
	// For loop: old = s.head.Load(), n.next = old.
	// If s.head.CompareAndSwap(old, n), break.
	panic("not implemented")
}
