package methodval

import "testing"

func TestBoundEarly(t *testing.T) {
	f := BoundEarly(7)
	if got := f(); got != 7 {
		t.Errorf("=%d want 7 (method value should have copied the early receiver)", got)
	}
}
