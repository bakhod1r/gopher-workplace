package tosetgen

import "testing"

func TestToSet(t *testing.T) {
	got := ToSet([]int{1, 1, 2})
	if len(got) != 2 {
		t.Fatalf("ToSet([]int{1, 1, 2}) has %d keys, want 2", len(got))
	}
	for _, k := range []int{1, 2} {
		if _, ok := got[k]; !ok {
			t.Errorf("key %v missing from set", k)
		}
	}
	if s := ToSet([]string{"a"}); len(s) != 1 {
		t.Errorf(`ToSet([]string{"a"}) has %d keys, want 1`, len(s))
	}
	empty := ToSet([]int{})
	if empty == nil {
		t.Error("ToSet([]int{}) = nil, want an empty map")
	}
	if len(empty) != 0 {
		t.Errorf("ToSet([]int{}) has %d keys, want 0", len(empty))
	}
}
