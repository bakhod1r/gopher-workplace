package toggler

import "testing"

func TestToggle(t *testing.T) {
	s := &Switch{}
	if s.State() {
		t.Error("new Switch should be off")
	}
	s.Toggle()
	if !s.State() {
		t.Error("after one Toggle should be on")
	}
	s.Toggle()
	if s.State() {
		t.Error("after two Toggles should be off")
	}
}

func TestToggleAll(t *testing.T) {
	a, b := &Switch{}, &Switch{On: true}
	if got := ToggleAll([]Toggler{a, b}); got != 1 {
		t.Errorf("ToggleAll = %d, want 1", got)
	}
	if !a.On {
		t.Error("a should be on")
	}
	if b.On {
		t.Error("b should be off")
	}
	if got := ToggleAll(nil); got != 0 {
		t.Errorf("ToggleAll(nil) = %d, want 0", got)
	}
}
