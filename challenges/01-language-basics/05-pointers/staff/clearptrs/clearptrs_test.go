package clearptrs

import "testing"

func TestClear(t *testing.T) {
	a, b := 1, 2
	s := []*int{&a, &b}
	out := Clear(s)
	if len(out) != 0 {
		t.Fatalf("len=%d want 0", len(out))
	}
	// backing array must be cleared of references
	full := s[:2]
	if full[0] != nil || full[1] != nil {
		t.Errorf("backing array still holds pointers (retention leak)")
	}
}
