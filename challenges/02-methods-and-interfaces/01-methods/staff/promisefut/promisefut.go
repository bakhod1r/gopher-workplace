// Package promisefut — Gopher Workplace challenge.
package promisefut

// Future is a value that becomes available later. Every Get blocks until the
// future is completed and then returns the same value.
type Future struct {
	done chan struct{}
	val  int
}

// NewFuture returns an uncompleted Future.
func NewFuture() *Future {
	return &Future{done: make(chan struct{})}
}

// Complete resolves the future with val and unblocks every waiter.
func (f *Future) Complete(val int) {
	// TODO(candidate): store val, then close f.done.
	panic("not implemented")
}

// Get blocks until the future is completed and returns its value.
func (f *Future) Get() int {
	// TODO(candidate): wait on f.done, then return f.val.
	panic("not implemented")
}

// IsDone reports whether the future has been completed, without blocking.
func (f *Future) IsDone() bool {
	select {
	case <-f.done:
		return true
	default:
		return false
	}
}
