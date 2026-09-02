// Package statemach — Gopher Workplace challenge.
package statemach

// State handles events and names itself.
type State interface {
	Next(event string) (State, bool)
	Name() string
}

// Pending is the initial state.
type Pending struct{}

// Name returns "pending".
func (p Pending) Name() string {
	// TODO(candidate): "pending".
	panic("not implemented")
}

// Next accepts "ship" (=> Shipped) and "cancel" (stays, reports false).
//
// Examples:
//
//	Pending{}.Next("ship") => Shipped{}, true
func (p Pending) Next(event string) (State, bool) {
	// TODO(candidate): handle "ship"; everything else stays put.
	panic("not implemented")
}

// Shipped is in transit.
type Shipped struct{}

// Name returns "shipped".
func (s Shipped) Name() string {
	// TODO(candidate): "shipped".
	panic("not implemented")
}

// Next accepts "deliver" (=> Delivered).
func (s Shipped) Next(event string) (State, bool) {
	// TODO(candidate): handle "deliver".
	panic("not implemented")
}

// Delivered is terminal.
type Delivered struct{}

// Name returns "delivered".
func (d Delivered) Name() string {
	// TODO(candidate): "delivered".
	panic("not implemented")
}

// Next accepts nothing.
func (d Delivered) Next(event string) (State, bool) {
	// TODO(candidate): terminal state.
	panic("not implemented")
}

// Run applies every event and returns the final state name and how many
// events were accepted.
func Run(s State, events []string) (string, int) {
	// TODO(candidate): fold the events through the states.
	panic("not implemented")
}
