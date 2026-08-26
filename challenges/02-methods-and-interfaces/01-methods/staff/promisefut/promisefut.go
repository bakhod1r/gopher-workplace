// Package promisefut — Gopher Workplace challenge.
package promisefut

// Future represents an eventual result.
type Future struct {
	ch chan int
}

func NewFuture() *Future {
	return &Future{ch: make(chan int, 1)}
}

// Complete resolves the future.
func (f *Future) Complete(val int) {
	// TODO(candidate): send val to ch, then close it
	panic("not implemented")
}

// Get blocks until resolved.
func (f *Future) Get() int {
	// TODO(candidate): return <-ch
	panic("not implemented")
}
