package memoizegen

import "testing"

func TestMemoizeCaches(t *testing.T) {
	calls := 0
	m := Memoize(func(n int) int { calls++; return n * 2 })
	if got := m(2); got != 4 {
		t.Errorf("m(2) = %v, want 4", got)
	}
	if got := m(2); got != 4 {
		t.Errorf("m(2) = %v, want 4", got)
	}
	if calls != 1 {
		t.Errorf("underlying function called %d times, want 1", calls)
	}
}

func TestMemoizeCachesZeroValues(t *testing.T) {
	calls := 0
	m := Memoize(func(n int) int { calls++; return 0 })
	m(1)
	m(1)
	if calls != 1 {
		t.Errorf("called %d times for a cached zero, want 1", calls)
	}
}

func TestMemoizeDistinctKeys(t *testing.T) {
	calls := 0
	m := Memoize(func(s string) int { calls++; return len(s) })
	m("a")
	m("bb")
	m("a")
	if calls != 2 {
		t.Errorf("called %d times, want 2", calls)
	}
	if got := m("bb"); got != 2 {
		t.Errorf(`m("bb") = %v, want 2`, got)
	}
}
