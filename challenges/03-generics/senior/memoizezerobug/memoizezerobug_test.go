package memoizezerobug

import "testing"

func TestMemoizeCaches(t *testing.T) {
	calls := 0
	f := Memoize(func(k int) int { calls++; return k * 2 })
	if got := f(3); got != 6 {
		t.Errorf("f(3) = %d, want 6", got)
	}
	f(3)
	f(3)
	if calls != 1 {
		t.Errorf("underlying called %d times, want 1", calls)
	}
}

func TestMemoizeDistinctKeys(t *testing.T) {
	calls := 0
	f := Memoize(func(k int) int { calls++; return k })
	f(1)
	f(2)
	if calls != 2 {
		t.Errorf("underlying called %d times, want 2", calls)
	}
}
