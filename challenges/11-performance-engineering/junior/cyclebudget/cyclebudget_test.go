package cyclebudget

import "testing"

func TestCyclesPerOp(t *testing.T) {
	cases := []struct {
		ns, ghz, want float64
	}{
		{10, 3, 30},
		{1, 1, 1},
		{0.5, 4, 2},
		{0, 3, 0},
		{10, 0, 0},
		{-5, 3, 0},
	}
	for _, c := range cases {
		if got := CyclesPerOp(c.ns, c.ghz); got != c.want {
			t.Errorf("CyclesPerOp(%v, %v) = %v, want %v", c.ns, c.ghz, got, c.want)
		}
	}
}

func TestVerdict(t *testing.T) {
	cases := []struct {
		cycles float64
		want   string
	}{
		{0, "register"},
		{9.9, "register"},
		{10, "l1"},
		{30, "l1"},
		{99, "l1"},
		{100, "memory"},
		{999, "memory"},
		{1000, "syscall"},
		{50000, "syscall"},
		{-1, "register"},
	}
	for _, c := range cases {
		if got := Verdict(c.cycles); got != c.want {
			t.Errorf("Verdict(%v) = %q, want %q", c.cycles, got, c.want)
		}
	}
}
