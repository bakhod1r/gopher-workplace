// Package fluxpattern — Gopher Workplace challenge.
package fluxpattern

// Store holds state.
type Store struct {
	Count int
}

// Dispatch processes an action string to mutate state.
func (s *Store) Dispatch(action string) {
	// TODO(candidate): if "INC", Count++. if "DEC", Count--.
	panic("not implemented")
}
