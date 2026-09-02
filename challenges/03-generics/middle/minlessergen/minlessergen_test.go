package minlessergen

import "testing"

func TestMinOf(t *testing.T) {
	got, ok := MinOf([]Version{{3}, {1}, {2}})
	if !ok || got.N != 1 {
		t.Errorf("MinOf = %+v, %v, want {1}, true", got, ok)
	}
}

func TestMinOfTieKeepsEarlier(t *testing.T) {
	vs := []Version{{2}, {2}}
	got, _ := MinOf(vs)
	if got != vs[0] {
		t.Errorf("MinOf = %+v, want the first element", got)
	}
}

func TestMinOfEmpty(t *testing.T) {
	got, ok := MinOf([]Version{})
	if ok || got.N != 0 {
		t.Errorf("MinOf(empty) = %+v, %v, want zero, false", got, ok)
	}
}

func TestMinOfAllIncreasing(t *testing.T) {
	got, _ := MinOf([]Version{{1}, {2}, {3}})
	if got.N != 1 {
		t.Errorf("MinOf = %+v, want {1}", got)
	}
}
