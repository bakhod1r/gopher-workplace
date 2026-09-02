// Package panicunwind — Gopher Workplace challenge.
package panicunwind

import "fmt"

// Task is a unit of work that may panic.
type Task interface {
	Run() error
}

// TaskFunc adapts a function to Task.
type TaskFunc func() error

// Run calls the underlying function.
func (f TaskFunc) Run() error { return f() }

// SafeRun runs t, converting a panic into an error.
//
// A task that returns an error normally has that error passed through
// unchanged.
//
// Examples:
//
//	a task panicking with "boom" => an error mentioning boom
func SafeRun(t Task) (err error) {
	// TODO(candidate): deferred recover assigning to the named result.
	panic("not implemented")
}

// RunAll runs every task, returning one result per task in order.
func RunAll(ts []Task) []error {
	// TODO(candidate): one SafeRun per task.
	panic("not implemented")
}

// Order appends a marker in each deferred call and returns the order they
// ran in, even though the function panics partway through.
func Order() []string {
	// TODO(candidate): three defers plus a panic; return the observed order.
	panic("not implemented")
}

var _ = fmt.Errorf
