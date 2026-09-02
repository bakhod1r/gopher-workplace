// Package pipelinegen — Gopher Workplace challenge.
package pipelinegen

// Pipeline returns a function applying every stage in order.
// With no stages it returns its input unchanged.
func Pipeline[T any](stages ...func(T) T) func(T) T {
	// TODO(candidate): return a closure applying every stage in order.
	panic("not implemented")
}
