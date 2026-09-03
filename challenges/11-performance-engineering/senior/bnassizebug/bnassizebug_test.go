package bnassizebug

import "testing"

func TestWork(t *testing.T) {
	if got := Work(3, 10); got != 30 {
		t.Errorf("Work(3, 10) = %d, want 30", got)
	}
	if got := Work(1, 10); got != 10 {
		t.Errorf("Work(1, 10) = %d, want 10", got)
	}
	if got := Work(10, 1); got != 10 {
		t.Errorf("Work(10, 1) = %d, want 10", got)
	}
}

func TestWorkGuards(t *testing.T) {
	for _, c := range []struct{ n, size int }{{0, 10}, {-1, 10}, {3, 0}, {3, -2}} {
		if got := Work(c.n, c.size); got != 0 {
			t.Errorf("Work(%d, %d) = %d, want 0", c.n, c.size, got)
		}
	}
}

func TestPerOpDoesNotDependOnN(t *testing.T) {
	want := PerOp(1, 10)
	for _, n := range []int{2, 10, 1000, 1_000_000} {
		if got := PerOp(n, 10); got != want {
			t.Fatalf("PerOp(%d, 10) = %d, want %d — the input size must not scale with b.N", n, got, want)
		}
	}
	if want != 10 {
		t.Errorf("PerOp = %d, want 10", want)
	}
}

func TestWorkScalesWithSize(t *testing.T) {
	if Work(5, 20) != 2*Work(5, 10) {
		t.Errorf("doubling the input size did not double the work: %d vs %d", Work(5, 20), Work(5, 10))
	}
}
