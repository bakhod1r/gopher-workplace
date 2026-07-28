package recoverdirect

import "testing"

func TestGuard(t *testing.T) {
	if Guard(func() {}) {
		t.Errorf("normal call should return false")
	}
	if !Guard(func() { panic("boom") }) {
		t.Errorf("panicking call should be recovered -> true")
	}
}
