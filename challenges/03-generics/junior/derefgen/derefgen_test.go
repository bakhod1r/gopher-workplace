package derefgen

import "testing"

func TestDeref(t *testing.T) {
	if got := Deref(Ptr(7), 0); got != 7 {
		t.Errorf("Deref(Ptr(7), 0) = %v, want 7", got)
	}
	if got := Deref(Ptr("a"), "z"); got != "a" {
		t.Errorf(`Deref(Ptr("a"), "z") = %q, want "a"`, got)
	}
	if got := Deref((*int)(nil), 0); got != 0 {
		t.Errorf("Deref(nil, 0) = %v, want 0", got)
	}
	if got := Deref((*string)(nil), "def"); got != "def" {
		t.Errorf(`Deref(nil, "def") = %q, want "def"`, got)
	}
	if got := Deref(Ptr(0), 9); got != 0 {
		t.Errorf("Deref(Ptr(0), 9) = %v, want 0", got)
	}
}
