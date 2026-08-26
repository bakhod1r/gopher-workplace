// Package mementopatt — Gopher Workplace challenge.
package mementopatt

// Editor holds text.
type Editor struct {
	Text string
}

// Memento is an opaque snapshot of state.
type Memento struct {
	state string
}

// Save returns a Memento representing current state.
func (e *Editor) Save() Memento {
	// TODO(candidate): return a Memento with current Text.
	panic("not implemented")
}

// Restore overwrites state with the Memento.
func (e *Editor) Restore(m Memento) {
	// TODO(candidate): set Text to m.state.
	panic("not implemented")
}
