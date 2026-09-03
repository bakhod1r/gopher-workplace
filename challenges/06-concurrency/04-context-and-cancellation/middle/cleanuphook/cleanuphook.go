// Package cleanuphook — Gopher Workplace challenge.
package cleanuphook

import "context"

// Hub tracks the teardown of one connection.
type Hub struct {
	closed chan struct{}
	stop   func() bool
}

// Register attaches a teardown to ctx: when ctx finishes, the connection is
// torn down and Wait returns "context". The returned Hub can Release the
// connection first, which stops the hook from running at all.
//
// Examples:
//
//	h := Register(ctx); cancel(); h.Wait()   => "context"
//	h := Register(ctx); h.Release()          => true
//	h := Register(cancelled ctx); h.Release() => false
func Register(ctx context.Context) *Hub {
	// TODO(candidate): use context.AfterFunc; keep its stop func on the Hub.
	panic("not implemented")
}

// Release tears the connection down early and reports whether it beat the
// context — true when the hook was stopped before it ran.
func (h *Hub) Release() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Wait blocks until the connection is torn down and reports who did it:
// "context" or "release".
func (h *Hub) Wait() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
