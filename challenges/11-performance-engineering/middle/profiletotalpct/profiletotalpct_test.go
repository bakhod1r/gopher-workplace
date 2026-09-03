package profiletotalpct

import (
	"reflect"
	"testing"
)

func TestTop(t *testing.T) {
	got := Top(map[string]int64{"a": 3, "b": 1})
	want := []Row{
		{"a", 3, 75, 75},
		{"b", 1, 25, 100},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Top = %v, want %v", got, want)
	}
}

func TestTopRounding(t *testing.T) {
	got := Top(map[string]int64{"a": 1, "b": 1, "c": 1})
	if len(got) != 3 {
		t.Fatalf("Top = %v, want 3 rows", got)
	}
	if got[0].Pct != 33.33 || got[1].Pct != 33.33 || got[2].Pct != 33.33 {
		t.Errorf("Pct values = %v, want 33.33 each", got)
	}
	if got[2].CumPct != 100 {
		t.Errorf("final CumPct = %v, want 100", got[2].CumPct)
	}
}

func TestTopOrderingAndTies(t *testing.T) {
	got := Top(map[string]int64{"z": 5, "a": 5, "m": 10})
	names := []string{got[0].Func, got[1].Func, got[2].Func}
	if !reflect.DeepEqual(names, []string{"m", "a", "z"}) {
		t.Errorf("order = %v, want [m a z]", names)
	}
}

func TestTopDropsNonPositive(t *testing.T) {
	got := Top(map[string]int64{"a": 4, "b": 0, "c": -1})
	want := []Row{{"a", 4, 100, 100}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Top = %v, want %v", got, want)
	}
}

func TestTopEmpty(t *testing.T) {
	for _, in := range []map[string]int64{nil, {"a": 0}} {
		got := Top(in)
		if got == nil || len(got) != 0 {
			t.Errorf("Top(%v) = %v, want empty non-nil slice", in, got)
		}
	}
}

func TestCoveringCount(t *testing.T) {
	flat := map[string]int64{"a": 50, "b": 30, "c": 20}
	cases := []struct {
		pct  float64
		want int
	}{
		{0, 0},
		{-5, 0},
		{50, 1},
		{51, 2},
		{80, 2},
		{100, 3},
		{150, 3},
	}
	for _, c := range cases {
		if got := CoveringCount(flat, c.pct); got != c.want {
			t.Errorf("CoveringCount(%v) = %d, want %d", c.pct, got, c.want)
		}
	}
}

func TestCoveringCountEmpty(t *testing.T) {
	if got := CoveringCount(nil, 50); got != 0 {
		t.Errorf("CoveringCount(nil) = %d, want 0", got)
	}
}
