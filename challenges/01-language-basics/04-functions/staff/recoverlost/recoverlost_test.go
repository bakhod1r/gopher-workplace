package recoverlost

import "testing"

func TestSafe(t *testing.T) {
	if err := Safe(func() { panic("boom") }); err == nil {
		t.Errorf("panic should yield a non-nil error")
	}
}
