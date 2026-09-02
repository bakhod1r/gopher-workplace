// Package workerpool — Gopher Workplace challenge.
package workerpool

// Task is one unit of work.
type Task interface {
	Run() int
}

// SquareTask squares its value.
type SquareTask struct {
	N int
}

// Run returns N*N.
func (s SquareTask) Run() int {
	// TODO(candidate): square it.
	panic("not implemented")
}

// RunAll runs every task using at most workers goroutines and returns the
// results in input order.
//
// Examples:
//
//	RunAll([]Task{SquareTask{2}, SquareTask{3}}, 2) => [4 9]
func RunAll(tasks []Task, workers int) []int {
	// TODO(candidate): bounded pool, index-addressed results.
	panic("not implemented")
}
