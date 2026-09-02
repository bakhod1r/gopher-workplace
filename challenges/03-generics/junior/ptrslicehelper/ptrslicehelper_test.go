package ptrslicehelper

import "testing"

func TestPtrsTo(t *testing.T) {
	got := PtrsTo([]int{1, 2, 3})
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for i, want := range []int{1, 2, 3} {
		if *got[i] != want {
			t.Errorf("*got[%d] = %v, want %v", i, *got[i], want)
		}
	}
}

func TestPtrsToAreDistinct(t *testing.T) {
	got := PtrsTo([]int{1, 2})
	if got[0] == got[1] {
		t.Error("both pointers address the same variable, want distinct addresses")
	}
	*got[0] = 99
	if *got[1] != 2 {
		t.Errorf("*got[1] = %v, want 2", *got[1])
	}
}

func TestPtrsToDoesNotAliasInput(t *testing.T) {
	in := []int{1}
	got := PtrsTo(in)
	*got[0] = 99
	if in[0] != 1 {
		t.Errorf("input mutated: in[0] = %v, want 1", in[0])
	}
}

func TestPtrsToEmpty(t *testing.T) {
	got := PtrsTo([]int{})
	if got == nil || len(got) != 0 {
		t.Errorf("PtrsTo([]) = %v, want an empty non-nil slice", got)
	}
}
