package preallocgen

import "testing"

func TestBuild(t *testing.T) {
	got := Build(3, func(i int) int { return i * i })
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for i, want := range []int{0, 1, 4} {
		if got[i] != want {
			t.Errorf("got[%d] = %v, want %v", i, got[i], want)
		}
	}
}

func TestBuildAllocatesExactly(t *testing.T) {
	got := Build(5, func(i int) string { return "x" })
	if cap(got) != 5 {
		t.Errorf("cap = %d, want exactly 5 (allocate once)", cap(got))
	}
	if len(got) != 5 {
		t.Errorf("len = %d, want 5", len(got))
	}
}

func TestBuildEdges(t *testing.T) {
	if got := Build(0, func(i int) int { return i }); len(got) != 0 {
		t.Errorf("Build(0) = %v, want []", got)
	}
	if got := Build(-1, func(i int) int { return i }); len(got) != 0 {
		t.Errorf("Build(-1) = %v, want [] (must not panic)", got)
	}
}
