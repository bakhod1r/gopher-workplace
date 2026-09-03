package quantileinterp

import (
	"math"
	"reflect"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestInterpolated(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	cases := []struct{ p, want float64 }{
		{0, 1},
		{50, 2.5},
		{100, 4},
		{25, 1.75},
		{75, 3.25},
	}
	for _, c := range cases {
		if got := Interpolated(s, c.p); !near(got, c.want) {
			t.Errorf("Interpolated(%v) = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestInterpolatedInventsValues(t *testing.T) {
	// 2.5 was never observed; that is the point of interpolation.
	got := Interpolated([]float64{1, 2, 3, 4}, 50)
	for _, v := range []float64{1, 2, 3, 4} {
		if got == v {
			t.Fatalf("Interpolated returned an observed sample %v", got)
		}
	}
}

func TestInterpolatedSortsACopy(t *testing.T) {
	in := []float64{4, 1, 3, 2}
	before := append([]float64(nil), in...)
	if got := Interpolated(in, 50); !near(got, 2.5) {
		t.Errorf("Interpolated = %v, want 2.5", got)
	}
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input was sorted in place: %v, want %v", in, before)
	}
}

func TestInterpolatedEdgeCases(t *testing.T) {
	if got := Interpolated(nil, 50); got != 0 {
		t.Errorf("Interpolated(nil) = %v, want 0", got)
	}
	if got := Interpolated([]float64{7}, 50); !near(got, 7) {
		t.Errorf("Interpolated = %v, want 7", got)
	}
	if got := Interpolated([]float64{1, 2}, 150); !near(got, 2) {
		t.Errorf("Interpolated(150) = %v, want 2", got)
	}
	if got := Interpolated([]float64{1, 2}, -5); !near(got, 1) {
		t.Errorf("Interpolated(-5) = %v, want 1", got)
	}
}

func TestNearestRank(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	if got := NearestRank(s, 50); got != 2 {
		t.Errorf("NearestRank = %v, want 2", got)
	}
	if got := NearestRank(s, 75); got != 3 {
		t.Errorf("NearestRank = %v, want 3", got)
	}
	if got := NearestRank(nil, 50); got != 0 {
		t.Errorf("NearestRank(nil) = %v, want 0", got)
	}
}

func TestTheTwoDefinitionsDisagree(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	if NearestRank(s, 50) == Interpolated(s, 50) {
		t.Error("the two definitions agreed on an even-sized sample; one of them is wrong")
	}
}
