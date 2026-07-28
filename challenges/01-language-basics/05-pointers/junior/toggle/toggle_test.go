package toggle

import "testing"

func TestToggle(t *testing.T) {
	b := false
	Toggle(&b)
	if !b {
		t.Errorf("want true")
	}
	Toggle(&b)
	if b {
		t.Errorf("want false")
	}
}
