package lesserselfbug

import "testing"

type ver struct {
	maj, min int
}

func (v ver) Less(o ver) bool {
	if v.maj != o.maj {
		return v.maj < o.maj
	}
	return v.min < o.min
}

func TestMinOfMiddle(t *testing.T) {
	got, ok := MinOf([]ver{{1, 5}, {1, 2}, {2, 0}})
	if !ok || got != (ver{1, 2}) {
		t.Errorf("MinOf = %v, %v; want {1 2}, true", got, ok)
	}
}

func TestMinOfFirstIsSmallest(t *testing.T) {
	got, ok := MinOf([]ver{{0, 1}, {9, 9}})
	if !ok || got != (ver{0, 1}) {
		t.Errorf("MinOf = %v, %v; want {0 1}, true", got, ok)
	}
}

func TestMinOfSingle(t *testing.T) {
	got, ok := MinOf([]ver{{4, 4}})
	if !ok || got != (ver{4, 4}) {
		t.Errorf("MinOf = %v, %v; want {4 4}, true", got, ok)
	}
}

func TestMinOfEmpty(t *testing.T) {
	if _, ok := MinOf([]ver{}); ok {
		t.Errorf("MinOf ok = true, want false")
	}
}
