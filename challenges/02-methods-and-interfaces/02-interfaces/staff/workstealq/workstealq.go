// Package workstealq — Gopher Workplace challenge.
package workstealq

import "sync"

// Deque is a work-stealing double-ended queue.
type Deque struct {
	mu     sync.Mutex
	items  []int
	top    int // index of the next steal
	bottom int // index one past the last push
}

// Stealer exposes the thief-facing end.
type Stealer interface {
	Steal() (int, bool)
}

// NewDeque returns an empty deque.
func NewDeque() *Deque {
	return &Deque{items: make([]int, 0, 16)}
}

// Push adds an item at the bottom. Only the owner calls it.
//
// Examples:
//
//	Push(1); Push(2); Pop() => 2
func (d *Deque) Push(v int) {
	// TODO(candidate): append at the bottom.
	panic("not implemented")
}

// Pop removes the most recently pushed item. Only the owner calls it.
func (d *Deque) Pop() (int, bool) {
	// TODO(candidate): take from the bottom.
	panic("not implemented")
}

// Steal removes the oldest item. Any goroutine may call it.
func (d *Deque) Steal() (int, bool) {
	// TODO(candidate): take from the top.
	panic("not implemented")
}

// Len returns how many items are queued.
func (d *Deque) Len() int {
	// TODO(candidate): bottom minus top.
	panic("not implemented")
}
