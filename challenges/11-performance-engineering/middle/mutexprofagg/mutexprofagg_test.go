package mutexprofagg

import (
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	count, delay, ok := Scale(3, 300, 5)
	if !ok || count != 15 || delay != 1500 {
		t.Errorf("Scale = %d, %d, %v, want 15, 1500, true", count, delay, ok)
	}
	count, delay, ok = Scale(3, 300, 1)
	if !ok || count != 3 || delay != 300 {
		t.Errorf("Scale = %d, %d, %v, want 3, 300, true", count, delay, ok)
	}
}

func TestScaleReportsAnOffProfile(t *testing.T) {
	for _, f := range []int{0, -1} {
		if _, _, ok := Scale(3, 300, f); ok {
			t.Errorf("Scale(_, _, %d) reported ok; a disabled profile cannot be scaled", f)
		}
	}
}

func TestEstimate(t *testing.T) {
	got := Estimate([]Contention{{"a", 1, 100}, {"b", 2, 40}, {"a", 1, 100}}, 5)
	want := map[string]int64{"a": 1000, "b": 200}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Estimate = %v, want %v", got, want)
	}
}

func TestEstimateWithTheProfileOff(t *testing.T) {
	got := Estimate([]Contention{{"a", 1, 100}}, 0)
	if got == nil || len(got) != 0 {
		t.Errorf("Estimate = %v, want empty non-nil map — nothing can be estimated from a disabled profile", got)
	}
}

func TestEstimateDropsJunk(t *testing.T) {
	got := Estimate([]Contention{{"a", 0, 100}, {"b", 1, -5}}, 1)
	if got == nil || len(got) != 0 {
		t.Errorf("Estimate = %v, want empty non-nil map", got)
	}
}

func TestConfidence(t *testing.T) {
	cases := []struct {
		samples int64
		want    string
	}{
		{0, "low"},
		{9, "low"},
		{10, "medium"},
		{99, "medium"},
		{100, "high"},
		{100000, "high"},
		{-5, "low"},
	}
	for _, c := range cases {
		if got := Confidence(c.samples); got != c.want {
			t.Errorf("Confidence(%d) = %q, want %q", c.samples, got, c.want)
		}
	}
}
