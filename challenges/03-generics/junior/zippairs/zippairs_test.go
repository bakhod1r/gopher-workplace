package zippairs

import "testing"

func TestZip(t *testing.T) {
	got := Zip([]int{1, 2}, []string{"a", "b"})
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].First != 1 || got[0].Second != "a" {
		t.Errorf("got[0] = %+v, want {1 a}", got[0])
	}
	if got[1].First != 2 || got[1].Second != "b" {
		t.Errorf("got[1] = %+v, want {2 b}", got[1])
	}
}

func TestZipStopsAtShorter(t *testing.T) {
	got := Zip([]int{1, 2, 3}, []string{"a"})
	if len(got) != 1 {
		t.Errorf("len = %d, want 1", len(got))
	}
	empty := Zip([]int{}, []string{"a"})
	if empty == nil || len(empty) != 0 {
		t.Errorf("Zip with an empty slice = %v, want an empty non-nil result", empty)
	}
}
