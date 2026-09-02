package signgen

import "testing"

func TestSign(t *testing.T) {
	if got := Sign(-2); got != -1 {
		t.Errorf("Sign(-2) = %v, want -1", got)
	}
	if got := Sign(0); got != 0 {
		t.Errorf("Sign(0) = %v, want 0", got)
	}
	if got := Sign(7); got != 1 {
		t.Errorf("Sign(7) = %v, want 1", got)
	}
	if got := Sign(1.5); got != 1 {
		t.Errorf("Sign(1.5) = %v, want 1", got)
	}
	if got := Sign(-0.5); got != -1 {
		t.Errorf("Sign(-0.5) = %v, want -1", got)
	}
}
