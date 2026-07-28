package repanic

import "testing"

func TestRun(t *testing.T) {
	if !Run(func() {}) {
		t.Errorf("normal call should be true")
	}
	if Run(func() { panic(ErrSentinel) }) {
		t.Errorf("sentinel should be absorbed -> false")
	}
	defer func() {
		if r := recover(); r != "other" {
			t.Errorf("unknown panic should propagate, got %v", r)
		}
	}()
	Run(func() { panic("other") })
	t.Errorf("unknown panic did not propagate")
}
