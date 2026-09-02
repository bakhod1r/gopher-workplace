// Package ctxtree — Gopher Workplace challenge.
package ctxtree

import (
	"errors"
	"sync"
)

// ErrCancelled is the cancellation cause.
var ErrCancelled = errors.New("cancelled")

// Canceller is a cancellable scope.
type Canceller interface {
	Cancel()
	Done() <-chan struct{}
	Err() error
}

// Node is one scope in the cancellation tree.
type Node struct {
	mu        sync.Mutex
	done      chan struct{}
	cancelled bool
	children  []*Node
}

// NewRoot returns a live root scope.
func NewRoot() *Node {
	// TODO(candidate): a live node with an open done channel.
	panic("not implemented")
}

// Child returns a new scope under n. If n is already cancelled, the child is
// born cancelled.
//
// Examples:
//
//	Child of a cancelled node => already cancelled
func (n *Node) Child() *Node {
	// TODO(candidate): register the child, inheriting cancellation.
	panic("not implemented")
}

// Cancel cancels this node and its whole subtree. It is idempotent.
func (n *Node) Cancel() {
	// TODO(candidate): close once, then cancel the children.
	panic("not implemented")
}

// Done returns a channel closed when this node is cancelled.
func (n *Node) Done() <-chan struct{} {
	// TODO(candidate): return the done channel.
	panic("not implemented")
}

// Err returns ErrCancelled once cancelled, otherwise nil.
func (n *Node) Err() error {
	// TODO(candidate): report the cancellation state.
	panic("not implemented")
}
