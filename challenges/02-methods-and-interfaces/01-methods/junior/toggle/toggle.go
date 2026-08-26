// Package toggle — Gopher Workplace challenge.
package toggle

// Switch holds an on/off state.
type Switch struct {
	On bool
}

// Toggle flips the switch: on → off, off → on.
//
// Examples:
//
//	s := Switch{On: false}; s.Toggle() // s.On == true
//	s.Toggle()                          // s.On == false
func (s *Switch) Toggle() {
	// TODO(candidate): implement this.
	panic("not implemented")
}
