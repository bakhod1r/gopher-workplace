package comparableembed

import "testing"

func TestTally(t *testing.T) {
	got := Tally([]int{1, 1, 2})
	if got[1] != 2 || got[2] != 1 {
		t.Errorf("Tally = %v, want {1:2 2:1}", got)
	}
	if len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
}

func TestTallyStrings(t *testing.T) {
	got := Tally([]string{"a", "b", "a"})
	if got["a"] != 2 || got["b"] != 1 {
		t.Errorf("Tally = %v, want {a:2 b:1}", got)
	}
}

func TestTallyEmpty(t *testing.T) {
	got := Tally([]int{})
	if got == nil {
		t.Fatal("Tally(empty) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("Tally(empty) = %v, want {}", got)
	}
}
