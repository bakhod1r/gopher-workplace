package pipelinegen

import "testing"

func TestPipeline(t *testing.T) {
	double := func(n int) int { return n * 2 }
	inc := func(n int) int { return n + 1 }
	if got := Pipeline(double, inc)(3); got != 7 {
		t.Errorf("Pipeline(double, inc)(3) = %v, want 7", got)
	}
	if got := Pipeline(inc, double)(3); got != 8 {
		t.Errorf("Pipeline(inc, double)(3) = %v, want 8", got)
	}
	if got := Pipeline(double)(3); got != 6 {
		t.Errorf("Pipeline(double)(3) = %v, want 6", got)
	}
}

func TestPipelineEmptyIsIdentity(t *testing.T) {
	if got := Pipeline[int]()(3); got != 3 {
		t.Errorf("Pipeline()(3) = %v, want 3", got)
	}
	if got := Pipeline[string]()("a"); got != "a" {
		t.Errorf(`Pipeline()("a") = %q, want "a"`, got)
	}
}
