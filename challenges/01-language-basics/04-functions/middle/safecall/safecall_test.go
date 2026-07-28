package safecall

import "testing"

func TestSafeInvoke(t *testing.T) {
	if SafeInvoke(func() {}) {
		t.Errorf("normal func should return false")
	}
	if !SafeInvoke(func() { panic("boom") }) {
		t.Errorf("panicking func should return true")
	}
}
