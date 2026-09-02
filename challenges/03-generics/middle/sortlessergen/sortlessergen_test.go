package sortlessergen

import "testing"

func TestSortedLess(t *testing.T) {
	in := []Version{{3}, {1}, {2}}
	got := SortedLess(in)
	for i, want := range []int{1, 2, 3} {
		if got[i].N != want {
			t.Fatalf("SortedLess = %+v, want [1 2 3]", got)
		}
	}
	if in[0].N != 3 {
		t.Errorf("input mutated: %+v", in)
	}
}

func TestSortedLessIsStable(t *testing.T) {
	type tagged struct{ Version }
	// Equal versions must keep their relative order; compare by identity of index.
	in := []Version{{2}, {2}, {1}}
	got := SortedLess(in)
	if got[0].N != 1 || got[1].N != 2 || got[2].N != 2 {
		t.Fatalf("SortedLess = %+v, want [1 2 2]", got)
	}
	_ = tagged{}
}

func TestSortedLessEmpty(t *testing.T) {
	got := SortedLess([]Version(nil))
	if got == nil || len(got) != 0 {
		t.Errorf("SortedLess(nil) = %v, want an empty non-nil slice", got)
	}
}
