// Package lockfreestk — Gopher Workplace challenge.
package lockfreestk

import "sync/atomic"

// node is one stack entry.
type node struct {
	value int
	next  *node
}

// Stack is a lock-free LIFO stack.
type Stack struct {
	head atomic.Pointer[node]
}

// Pusher accepts values.
type Pusher interface {
	Push(v int)
}

// Push adds a value to the top of the stack.
//
// Examples:
//
//	Push(1); Push(2); Pop() => 2, true
func (s *Stack) Push(v int) {
	// TODO(candidate): CAS loop installing a new head.
	panic("not implemented")
}

// Pop removes and returns the top value.
func (s *Stack) Pop() (int, bool) {
	// TODO(candidate): CAS loop unlinking the head.
	panic("not implemented")
}

// Len counts the linked nodes.
func (s *Stack) Len() int {
	// TODO(candidate): walk the chain from the head.
	panic("not implemented")
}
