package lazygen

import "testing"

func TestLazyComputesOnce(t *testing.T) {
	calls := 0
	l := NewLazy(func() int { calls++; return 42 })
	if l.Done() {
		t.Error("Done() before Get() = true, want false")
	}
	if got := l.Get(); got != 42 {
		t.Errorf("Get() = %v, want 42", got)
	}
	if got := l.Get(); got != 42 {
		t.Errorf("Get() = %v, want 42", got)
	}
	if calls != 1 {
		t.Errorf("compute called %d times, want 1", calls)
	}
	if !l.Done() {
		t.Error("Done() after Get() = false, want true")
	}
}

func TestLazyZeroValueIsRemembered(t *testing.T) {
	calls := 0
	l := NewLazy(func() int { calls++; return 0 })
	l.Get()
	l.Get()
	if calls != 1 {
		t.Errorf("compute called %d times for a zero result, want 1", calls)
	}
}
