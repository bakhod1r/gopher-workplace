// Package toggler — Gopher Workplace challenge.
package toggler

// Toggler flips between on and off.
type Toggler interface {
	Toggle()
	State() bool
}

// Switch is one on/off control.
type Switch struct {
	On bool
}

// Toggle flips the switch.
//
// Examples:
//
//	s := &Switch{}; s.Toggle() => s.State() == true
func (s *Switch) Toggle() {
	// TODO(candidate): invert On.
	panic("not implemented")
}

// State reports whether the switch is on.
func (s *Switch) State() bool {
	// TODO(candidate): report On.
	panic("not implemented")
}

// ToggleAll toggles every element and returns how many are now on.
func ToggleAll(ts []Toggler) int {
	// TODO(candidate): toggle each, count the on states.
	panic("not implemented")
}
