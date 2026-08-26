package lazyinit

import "testing"

func TestLazy(t *testing.T) {
	calls := 0
	l := New(func() string {
		calls++
		return "heavy"
	})

	if calls != 0 {
		t.Error("init called early")
	}

	if got := l.String(); got != "heavy" || calls != 1 {
		t.Errorf("first got=%q, calls=%d", got, calls)
	}

	if got := l.String(); got != "heavy" || calls != 1 {
		t.Errorf("second got=%q, calls=%d", got, calls)
	}
}
