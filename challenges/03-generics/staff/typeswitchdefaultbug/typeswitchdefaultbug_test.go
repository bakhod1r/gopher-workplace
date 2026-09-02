package typeswitchdefaultbug

import "testing"

func TestNormalizeString(t *testing.T) {
	if got := Normalize(" Hi "); got != "hi" {
		t.Errorf("Normalize = %q, want \"hi\"", got)
	}
}

func TestNormalizeInt(t *testing.T) {
	if got := Normalize(42); got != 42 {
		t.Errorf("Normalize = %d, want 42", got)
	}
}

func TestNormalizeOtherTypes(t *testing.T) {
	if got := Normalize(true); got != true {
		t.Errorf("Normalize = %v, want true", got)
	}
	if got := Normalize(3.5); got != 3.5 {
		t.Errorf("Normalize = %v, want 3.5", got)
	}
	type pt struct{ X, Y int }
	if got := Normalize(pt{1, 2}); got != (pt{1, 2}) {
		t.Errorf("Normalize = %v, want {1 2}", got)
	}
}
