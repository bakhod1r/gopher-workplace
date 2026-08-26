package countmin

import "testing"

func TestCountMin(t *testing.T) {
	s := &Sketch{}
	s.Add("apple")
	s.Add("apple")
	s.Add("ape")

	if got := s.Count("apple"); got != 2 {
		t.Errorf("apple count = %d, want 2", got)
	}
	if got := s.Count("ape"); got != 1 {
		t.Errorf("ape count = %d, want 1", got)
	}
}
