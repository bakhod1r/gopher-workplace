// Package epollwrap — Gopher Workplace challenge.
package epollwrap

import "sort"

// Event bits, mirroring the epoll flags they are named after.
const (
	EventRead  uint32 = 1 << 0
	EventWrite uint32 = 1 << 1
)

// Epoll tracks which events each registered file descriptor cares about.
type Epoll struct {
	interest map[int]uint32
}

// New returns an empty Epoll instance.
func New() *Epoll {
	return &Epoll{interest: make(map[int]uint32)}
}

// Add registers fd with the given interest mask, replacing any previous mask.
func (e *Epoll) Add(fd int, events uint32) {
	e.interest[fd] = events
}

// Remove deregisters fd. Removing an unregistered fd is a no-op.
func (e *Epoll) Remove(fd int) {
	delete(e.interest, fd)
}

// Wait reports which registered descriptors have at least one event the caller
// asked for. ready maps a descriptor to the events the kernel reports for it;
// entries for descriptors that are not registered are ignored. The returned
// descriptors are sorted ascending.
func (e *Epoll) Wait(ready map[int]uint32) []int {
	// TODO(candidate): intersect each ready mask with the registered interest
	// mask, keep the fds with a non-zero intersection, return them sorted.
	_ = sort.Ints
	panic("not implemented")
}
