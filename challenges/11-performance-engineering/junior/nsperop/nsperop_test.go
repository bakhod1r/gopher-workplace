package nsperop

import "testing"

func TestNsPerOp(t *testing.T) {
	cases := []struct {
		ns    int64
		iters int
		want  int64
	}{
		{1000, 3, 333},
		{1000, 1, 1000},
		{0, 10, 0},
		{9, 10, 0},
	}
	for _, c := range cases {
		if got := NsPerOp(c.ns, c.iters); got != c.want {
			t.Errorf("NsPerOp(%d, %d) = %d, want %d", c.ns, c.iters, got, c.want)
		}
	}
}

func TestNsPerOpGuards(t *testing.T) {
	if got := NsPerOp(1000, 0); got != 0 {
		t.Errorf("NsPerOp(1000, 0) = %d, want 0", got)
	}
	if got := NsPerOp(-5, 10); got != 0 {
		t.Errorf("NsPerOp(-5, 10) = %d, want 0", got)
	}
}

func TestFaster(t *testing.T) {
	cases := []struct {
		base, cand int64
		pct        float64
		want       bool
	}{
		{100, 80, 20, true},  // exactly 20% faster counts
		{100, 81, 20, false}, // 19% is not enough
		{100, 50, 20, true},
		{100, 120, 20, false},
		{0, 1, 20, false},
		{-1, 1, 20, false},
	}
	for _, c := range cases {
		if got := Faster(c.base, c.cand, c.pct); got != c.want {
			t.Errorf("Faster(%d, %d, %v) = %v, want %v", c.base, c.cand, c.pct, got, c.want)
		}
	}
}
