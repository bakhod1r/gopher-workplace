// Package pipelinestg — Gopher Workplace challenge.
package pipelinestg

// Stage transforms one value; emit is false to drop it.
type Stage interface {
	Process(v int) (out int, emit bool)
}

// DoubleStage doubles every value.
type DoubleStage struct{}

// Process doubles v.
func (DoubleStage) Process(v int) (int, bool) {
	// TODO(candidate): double, always emit.
	panic("not implemented")
}

// DropOddStage passes even values through and drops odd ones.
type DropOddStage struct{}

// Process emits only even values.
func (DropOddStage) Process(v int) (int, bool) {
	// TODO(candidate): emit only even values.
	panic("not implemented")
}

// RunStage reads in, applies s, and returns the output channel.
// The output channel closes when in closes.
//
// Examples:
//
//	DoubleStage over [1 2] => [2 4]
func RunStage(in <-chan int, s Stage) <-chan int {
	// TODO(candidate): one goroutine, deferred close, range over in.
	panic("not implemented")
}
