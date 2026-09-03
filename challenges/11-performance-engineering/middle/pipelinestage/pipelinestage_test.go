package pipelinestage

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

func double(v int) int { return v * 2 }
func inc(v int) int    { return v + 1 }

func TestStage(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	out := Stage(in, double)
	var got []int
	for v := range out {
		got = append(got, v)
	}
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("Stage = %v, want [2 4 6]", got)
	}
}

func TestStageClosesItsOutput(t *testing.T) {
	in := make(chan int)
	close(in)
	out := Stage(in, double)
	if _, ok := <-out; ok {
		t.Error("output channel produced a value from a closed empty input")
	}
	if _, ok := <-out; ok {
		t.Error("output channel is not closed when the input closes")
	}
}

func TestRun(t *testing.T) {
	got := Run([]int{1, 2, 3}, []func(int) int{double, inc})
	if !reflect.DeepEqual(got, []int{3, 5, 7}) {
		t.Errorf("Run = %v, want [3 5 7]", got)
	}
}

func TestRunNoStagesIsIdentity(t *testing.T) {
	got := Run([]int{1, 2}, nil)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Run = %v, want [1 2]", got)
	}
}

func TestRunEmptyInput(t *testing.T) {
	got := Run(nil, []func(int) int{double})
	if got == nil || len(got) != 0 {
		t.Errorf("Run(nil) = %v, want empty non-nil slice", got)
	}
}

func TestRunPreservesOrder(t *testing.T) {
	values := make([]int, 500)
	for i := range values {
		values[i] = i
	}
	got := Run(values, []func(int) int{double, inc, double})
	for i, v := range got {
		if want := ((i*2)+1)*2; v != want {
			t.Fatalf("result[%d] = %d, want %d", i, v, want)
		}
	}
}

func TestRunLeavesNoGoroutinesBehind(t *testing.T) {
	before := runtime.NumGoroutine()
	for i := 0; i < 50; i++ {
		Run([]int{1, 2, 3}, []func(int) int{double, inc, double})
	}
	// Give any leaked goroutines a chance to be counted.
	time.Sleep(50 * time.Millisecond)
	runtime.GC()
	if after := runtime.NumGoroutine(); after > before+5 {
		t.Errorf("goroutines: %d before, %d after — stages must exit when their input closes", before, after)
	}
}
