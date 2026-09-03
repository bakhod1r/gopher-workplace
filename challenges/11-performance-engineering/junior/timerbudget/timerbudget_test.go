package timerbudget

import "testing"

func TestCost(t *testing.T) {
	cases := []struct {
		ns    int64
		calls int
		want  int64
	}{
		{50, 20, 1000},
		{1, 1, 1},
		{50, 0, 0},
		{0, 20, 0},
		{50, -3, 0},
		{-50, 20, 0},
	}
	for _, c := range cases {
		if got := Cost(c.ns, c.calls); got != c.want {
			t.Errorf("Cost(%d, %d) = %d, want %d", c.ns, c.calls, got, c.want)
		}
	}
}

func TestHeadroom(t *testing.T) {
	if got := Headroom(50, 20, 1500); got != 500 {
		t.Errorf("Headroom = %d, want 500", got)
	}
	if got := Headroom(50, 20, 1000); got != 0 {
		t.Errorf("Headroom = %d, want 0", got)
	}
	if got := Headroom(50, 20, 400); got != -600 {
		t.Errorf("Headroom = %d, want -600 — an overrun keeps its sign", got)
	}
}

func TestFits(t *testing.T) {
	if !Fits(50, 20, 1500) {
		t.Error("Fits = false, want true")
	}
	if !Fits(50, 20, 1000) {
		t.Error("Fits at exactly the budget = false, want true")
	}
	if Fits(50, 21, 1000) {
		t.Error("Fits over the budget = true, want false")
	}
}

func TestDeathByAThousandCalls(t *testing.T) {
	// A "fast" 200ns helper called 10000 times per request is 2ms.
	if Fits(200, 10000, 1_000_000) {
		t.Error("2ms of work reported as fitting a 1ms budget")
	}
	if got := Cost(200, 10000); got != 2_000_000 {
		t.Errorf("Cost = %d, want 2000000", got)
	}
}
