package regressionflag

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		percent, tol float64
		want         string
	}{
		{-20, 5, "improvement"},
		{20, 5, "regression"},
		{0, 5, "noise"},
		{3, 5, "noise"},
		{-3, 5, "noise"},
		{5, 5, "noise"},
		{-5, 5, "noise"},
		{5.1, 5, "regression"},
		{-5.1, 5, "improvement"},
	}
	for _, c := range cases {
		if got := Classify(c.percent, c.tol); got != c.want {
			t.Errorf("Classify(%v, %v) = %q, want %q", c.percent, c.tol, got, c.want)
		}
	}
}

func TestClassifyZeroTolerance(t *testing.T) {
	if got := Classify(0.1, 0); got != "regression" {
		t.Errorf("Classify(0.1, 0) = %q, want regression", got)
	}
	if got := Classify(0, 0); got != "noise" {
		t.Errorf("Classify(0, 0) = %q, want noise", got)
	}
	if got := Classify(1, -5); got != "regression" {
		t.Errorf("Classify(1, -5) = %q, want regression — a negative tolerance is 0", got)
	}
}

func TestFailing(t *testing.T) {
	if Failing([]float64{-10, 2}, 5) {
		t.Error("Failing = true, want false")
	}
	if !Failing([]float64{-10, 2, 30}, 5) {
		t.Error("Failing = false, want true — one regression fails the suite")
	}
	if Failing(nil, 5) {
		t.Error("Failing(nil) = true, want false")
	}
	if Failing([]float64{-99}, 5) {
		t.Error("an improvement was reported as failing")
	}
}
