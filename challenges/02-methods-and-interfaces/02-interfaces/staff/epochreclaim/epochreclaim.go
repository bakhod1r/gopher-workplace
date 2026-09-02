// Package epochreclaim — Gopher Workplace challenge.
package epochreclaim

import "sync"

// Freer is an object whose memory can be released.
type Freer interface {
	Free()
}

// Node is a retired object that records whether it was freed.
type Node struct {
	ID    int
	Freed bool
}

// Free marks the node released.
func (n *Node) Free() { n.Freed = true }

// retired is an object waiting for its grace period.
type retired struct {
	obj   Freer
	epoch uint64
}

// Epoch tracks readers and defers frees until it is safe.
type Epoch struct {
	mu      sync.Mutex
	current uint64
	active  map[uint64]int // epoch => active readers
	pending []retired
}

// NewEpoch returns an empty epoch tracker.
func NewEpoch() *Epoch {
	return &Epoch{active: make(map[uint64]int)}
}

// Enter registers a reader and returns the epoch it entered.
func (e *Epoch) Enter() uint64 {
	// TODO(candidate): count a reader in the current epoch.
	panic("not implemented")
}

// Exit deregisters a reader from the given epoch.
func (e *Epoch) Exit(epoch uint64) {
	// TODO(candidate): drop the reader count for that epoch.
	panic("not implemented")
}

// Retire marks an object unreachable; it is freed once its grace period ends.
//
// Retiring advances the epoch so later readers cannot observe the object.
//
// Examples:
//
//	Retire while a reader is inside => not freed until that reader exits
func (e *Epoch) Retire(obj Freer) {
	// TODO(candidate): record the object with the current epoch, then advance.
	panic("not implemented")
}

// Reclaim frees every retired object whose grace period has passed and
// returns how many were freed.
func (e *Epoch) Reclaim() int {
	// TODO(candidate): free objects with no reader at or before their epoch.
	panic("not implemented")
}

// Pending returns how many objects are waiting to be freed.
func (e *Epoch) Pending() int {
	// TODO(candidate): count the waiting objects.
	panic("not implemented")
}
