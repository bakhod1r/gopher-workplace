// Package statemachinegen — Gopher Workplace challenge.
package statemachinegen

// Machine is a finite state machine over states S and events E.
// Use NewMachine to create one.
type Machine[S comparable, E comparable] struct {
	state S
	table map[S]map[E]S
}

// NewMachine returns a machine starting in the given state.
func NewMachine[S comparable, E comparable](start S) *Machine[S, E] {
	// TODO(candidate): store the start state and allocate the transition table.
	panic("not implemented")
}

// Allow records that event e moves the machine from a to b.
func (m *Machine[S, E]) Allow(a S, e E, b S) {
	// TODO(candidate): record the transition.
	panic("not implemented")
}

// Fire applies e, reporting whether the transition was allowed.
// A rejected event leaves the state unchanged.
func (m *Machine[S, E]) Fire(e E) bool {
	// TODO(candidate): apply the transition when one is defined.
	panic("not implemented")
}

// State returns the current state.
func (m *Machine[S, E]) State() S {
	// TODO(candidate): report the current state.
	panic("not implemented")
}
