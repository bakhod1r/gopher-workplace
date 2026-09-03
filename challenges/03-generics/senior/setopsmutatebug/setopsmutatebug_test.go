package setopsmutatebug

import "testing"

func set(vs ...int) map[int]struct{} {
	m := make(map[int]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	return m
}

func TestUnionContents(t *testing.T) {
	got := Union(set(1, 2), set(2, 3))
	for _, want := range []int{1, 2, 3} {
		if _, ok := got[want]; !ok {
			t.Errorf("union missing %d", want)
		}
	}
	if len(got) != 3 {
		t.Errorf("len(union) = %d, want 3", len(got))
	}
}

func TestUnionDoesNotMutateInputs(t *testing.T) {
	a, b := set(1, 2), set(2, 3)
	Union(a, b)
	if len(a) != 2 {
		t.Errorf("a mutated: len = %d, want 2", len(a))
	}
	if len(b) != 2 {
		t.Errorf("b mutated: len = %d, want 2", len(b))
	}
}
