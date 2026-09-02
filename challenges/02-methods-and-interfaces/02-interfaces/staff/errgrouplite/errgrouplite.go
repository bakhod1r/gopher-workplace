// Package errgrouplite — Gopher Workplace challenge.
package errgrouplite

import "sync"

// Task is one unit of work; it should stop early when cancel is closed.
type Task interface {
	Run(cancel <-chan struct{}) error
}

// TaskFunc adapts a function to Task.
type TaskFunc func(cancel <-chan struct{}) error

// Run calls the underlying function.
func (f TaskFunc) Run(cancel <-chan struct{}) error { return f(cancel) }

// Group runs tasks with bounded concurrency and first-error semantics.
type Group struct {
	Limit int

	initOnce sync.Once
	sem      chan struct{}
	wg       sync.WaitGroup

	errOnce sync.Once
	err     error
	cancel  chan struct{}
}

// NewGroup returns a group running at most limit tasks at once.
func NewGroup(limit int) *Group {
	if limit < 1 {
		limit = 1
	}
	return &Group{
		Limit:  limit,
		sem:    make(chan struct{}, limit),
		cancel: make(chan struct{}),
	}
}

// Go runs t in the group, blocking while the concurrency limit is reached.
//
// Examples:
//
//	three tasks, the second fails => Wait returns that error
func (g *Group) Go(t Task) {
	// TODO(candidate): acquire a slot, spawn, record the first error.
	panic("not implemented")
}

// Wait blocks until every task has finished and returns the first error.
func (g *Group) Wait() error {
	// TODO(candidate): wait, then report.
	panic("not implemented")
}

// Cancelled returns the group's cancellation channel.
func (g *Group) Cancelled() <-chan struct{} {
	// TODO(candidate): expose the cancel channel.
	panic("not implemented")
}
