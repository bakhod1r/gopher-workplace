package fluxpattern

import "testing"

func TestFlux(t *testing.T) {
	s := &Store{}
	s.Dispatch("INC")
	s.Dispatch("INC")
	s.Dispatch("DEC")

	if s.Count != 1 {
		t.Errorf("Count = %d, want 1", s.Count)
	}
}
