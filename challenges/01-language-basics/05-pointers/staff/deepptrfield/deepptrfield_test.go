package deepptrfield

import "testing"

func TestClone(t *testing.T) {
	x := 10
	b := &Box{P: &x}
	c := Clone(b)
	*c.P = 99
	if *b.P != 10 {
		t.Errorf("*b.P=%d want 10 (shared pointer field)", *b.P)
	}
}
