package p50p90p99

import (
	"reflect"
	"testing"
)

func TestPercentileNearestRank(t *testing.T) {
	s := []float64{1, 2, 3, 4}
	cases := []struct{ p, want float64 }{
		{50, 2},
		{25, 1},
		{75, 3},
		{100, 4},
		{0, 1},
		{1, 1},
	}
	for _, c := range cases {
		if got := Percentile(s, c.p); got != c.want {
			t.Errorf("Percentile(%v) = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestPercentileUnsortedInput(t *testing.T) {
	if got := Percentile([]float64{4, 1, 3, 2}, 50); got != 2 {
		t.Errorf("Percentile = %v, want 2", got)
	}
}

func TestPercentileDoesNotModifyInput(t *testing.T) {
	in := []float64{4, 1, 3, 2}
	before := append([]float64(nil), in...)
	Percentile(in, 50)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input was sorted in place: %v, want %v", in, before)
	}
}

func TestPercentileEdgeCases(t *testing.T) {
	if got := Percentile(nil, 50); got != 0 {
		t.Errorf("Percentile(nil) = %v, want 0", got)
	}
	if got := Percentile([]float64{7}, 99); got != 7 {
		t.Errorf("Percentile = %v, want 7", got)
	}
	if got := Percentile([]float64{1, 2}, 150); got != 2 {
		t.Errorf("Percentile(150) = %v, want 2", got)
	}
	if got := Percentile([]float64{1, 2}, -5); got != 1 {
		t.Errorf("Percentile(-5) = %v, want 1", got)
	}
}

func TestSummary(t *testing.T) {
	p50, p90, p99 := Summary([]float64{1, 2, 3, 4})
	if p50 != 2 || p90 != 4 || p99 != 4 {
		t.Errorf("Summary = %v, %v, %v, want 2, 4, 4", p50, p90, p99)
	}
}

func TestSummaryTailIsNotTheMean(t *testing.T) {
	// 99 fast requests and one very slow one.
	s := make([]float64, 0, 100)
	for i := 0; i < 99; i++ {
		s = append(s, 1)
	}
	s = append(s, 1000)
	p50, _, p99 := Summary(s)
	if p50 != 1 {
		t.Errorf("p50 = %v, want 1", p50)
	}
	if p99 != 1 {
		t.Errorf("p99 = %v, want 1 — rank 99 of 100 is still a fast request", p99)
	}
	if got := Percentile(s, 100); got != 1000 {
		t.Errorf("p100 = %v, want 1000", got)
	}
}
