package lazyzerobug

import "testing"

func TestLazyComputesOnce(t *testing.T) {
	calls := 0
	l := &Lazy[int]{Fn: func() int { calls++; return 7 }}
	if got := l.Get(); got != 7 {
		t.Errorf("Get = %d, want 7", got)
	}
	l.Get()
	if calls != 1 {
		t.Errorf("Fn called %d times, want 1", calls)
	}
}

func TestLazyCachesZeroResult(t *testing.T) {
	calls := 0
	l := &Lazy[int]{Fn: func() int { calls++; return 0 }}
	l.Get()
	l.Get()
	l.Get()
	if calls != 1 {
		t.Errorf("Fn called %d times for a zero result, want 1", calls)
	}
}
