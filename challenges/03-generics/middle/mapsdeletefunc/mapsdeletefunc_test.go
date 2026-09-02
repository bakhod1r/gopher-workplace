package mapsdeletefunc

import "testing"

func TestPrune(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Prune(m, func(k string, v int) bool { return v%2 == 0 })
	if len(got) != 2 || got["a"] != 1 || got["c"] != 3 {
		t.Errorf("Prune = %v, want {a:1 c:3}", got)
	}
	if len(m) != 3 {
		t.Errorf("input mutated: %v", m)
	}
}

func TestPruneByKey(t *testing.T) {
	m := map[string]int{"keep": 1, "drop": 2}
	got := Prune(m, func(k string, v int) bool { return k == "drop" })
	if len(got) != 1 || got["keep"] != 1 {
		t.Errorf("Prune = %v, want {keep:1}", got)
	}
}

func TestPruneNil(t *testing.T) {
	got := Prune(map[string]int(nil), func(string, int) bool { return true })
	if got == nil {
		t.Fatal("Prune(nil) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("Prune(nil) = %v, want {}", got)
	}
}
