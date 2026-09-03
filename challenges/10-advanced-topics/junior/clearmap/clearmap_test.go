package clearmap

import "testing"

func TestResetEmpties(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	Reset(m)
	if len(m) != 0 {
		t.Errorf("len(m) = %d, want 0", len(m))
	}
}

func TestResetIsVisibleToTheCaller(t *testing.T) {
	m := map[string]int{"a": 1}
	alias := m
	Reset(m)
	if len(alias) != 0 {
		t.Errorf("len(alias) = %d, want 0: the caller's map was not emptied", len(alias))
	}
	m["c"] = 3
	if alias["c"] != 3 {
		t.Error("the map was replaced, not emptied")
	}
}

func TestResetNilMap(t *testing.T) {
	Reset(nil)
}
