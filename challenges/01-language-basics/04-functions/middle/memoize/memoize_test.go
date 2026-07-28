package memoize

import "testing"

func TestMemoize(t *testing.T) {
	calls := 0
	slow := func(x int) int { calls++; return x * x }
	m := Memoize(slow)
	if m(4) != 16 || m(4) != 16 {
		t.Errorf("wrong result")
	}
	if calls != 1 {
		t.Errorf("f called %d times, want 1 (cached)", calls)
	}
	m(5)
	if calls != 2 {
		t.Errorf("new arg should call f: calls=%d", calls)
	}
}
