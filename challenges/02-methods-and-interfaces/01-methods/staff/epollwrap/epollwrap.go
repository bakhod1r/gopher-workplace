// Package epollwrap — Gopher Workplace challenge.
package epollwrap

type Epoll struct {
	Active bool
}

func (e *Epoll) Wait() bool {
	// TODO(candidate): return e.Active
	panic("not implemented")
}
