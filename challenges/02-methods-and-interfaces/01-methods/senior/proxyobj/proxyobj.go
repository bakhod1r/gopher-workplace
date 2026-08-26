// Package proxyobj — Gopher Workplace challenge.
package proxyobj

// Worker does expensive work.
type Worker struct {
	calls int
}

func (w *Worker) Do() string {
	w.calls++
	return "done"
}

// Proxy restricts access to the Worker.
type Proxy struct {
	w    *Worker
	role string
}

// Do only allows access if role is "admin". Otherwise returns "denied".
func (p *Proxy) Do() string {
	// TODO(candidate): check p.role, if "admin" call p.w.Do(), else "denied"
	panic("not implemented")
}
