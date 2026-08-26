package rwlockopt

import "testing"

func TestOptLock(t *testing.T) {
	o := &OptLock{}
	if got := o.IncrementIfZero(); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
	if got := o.IncrementIfZero(); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}
