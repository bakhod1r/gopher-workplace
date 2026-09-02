// Package workerpool — Gopher Workplace challenge.
package workerpool

// StopWorkers starts n image-resize workers that are idle, waiting on the pool
// context, then triggers the pool shutdown and waits for every worker to
// finish. It returns the stop reason each worker observed, indexed by worker
// number.
//
// Every element is context.Canceled, and the slice has exactly n elements.
//
// Examples:
//
//	StopWorkers(0)  => []
//	StopWorkers(1)  => [context.Canceled]
//	StopWorkers(3)  => [context.Canceled, context.Canceled, context.Canceled]
func StopWorkers(n int) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
