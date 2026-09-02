package pipelinerebuildbug

import (
	"reflect"
	"testing"
	"time"
)

var sink []int

func inc(n int) int    { return n + 1 }
func double(n int) int { return n * 2 }

func TestPipelineApplies(t *testing.T) {
	got := Pipeline([]int{1, 2}, inc, double)
	if !reflect.DeepEqual(got, []int{4, 6}) {
		t.Errorf("Pipeline = %v, want [4 6]", got)
	}
}

func TestPipelineDoesNotTouchTheInput(t *testing.T) {
	in := []int{1, 2}
	Pipeline(in, inc, double)
	if !reflect.DeepEqual(in, []int{1, 2}) {
		t.Errorf("input mutated: %v, want [1 2]", in)
	}
}

func TestPipelineAllocatesOneBuffer(t *testing.T) {
	in := make([]int, 4096)
	stages := []func(int) int{inc, inc, inc, inc, inc}
	allocs := testing.AllocsPerRun(5, func() { sink = Pipeline(in, stages...) })
	if allocs > 2 {
		t.Errorf("Pipeline allocated %.0f buffers for 5 stages, want at most 2", allocs)
	}
}

func TestPipelineScale(t *testing.T) {
	const n = 2_000_000
	in := make([]int, n)
	stages := []func(int) int{inc, inc, inc, inc, inc}
	start := time.Now()
	got := Pipeline(in, stages...)
	elapsed := time.Since(start)
	if len(got) != n || got[0] != 5 {
		t.Fatalf("Pipeline = len %d, first %d; want len %d, first 5", len(got), got[0], n)
	}
	if elapsed > 700*time.Millisecond {
		t.Errorf("5 stages over %d elements took %v, want under 700ms", n, elapsed)
	}
}
