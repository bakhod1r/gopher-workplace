package subbenchsharedstatebug

import (
	"reflect"
	"testing"
)

func TestRunSizeIsIndependent(t *testing.T) {
	var r Runner
	if got := r.RunSize(3); got != 3 {
		t.Errorf("RunSize(3) = %d, want 3", got)
	}
	if got := r.RunSize(3); got != 3 {
		t.Errorf("second RunSize(3) = %d, want 3 — the previous run must not carry over", got)
	}
}

func TestRunAll(t *testing.T) {
	got := RunAll([]int{1, 2, 3})
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("RunAll = %v, want [1 2 3]", got)
	}
}

func TestRunAllOrderDoesNotChangeResults(t *testing.T) {
	ascending := RunAll([]int{1, 10, 100})
	descending := RunAll([]int{100, 10, 1})
	if !reflect.DeepEqual(ascending, []int{1, 10, 100}) {
		t.Errorf("ascending = %v, want [1 10 100]", ascending)
	}
	if !reflect.DeepEqual(descending, []int{100, 10, 1}) {
		t.Errorf("descending = %v, want [100 10 1] — results must not depend on run order", descending)
	}
}

func TestRunSizeZero(t *testing.T) {
	var r Runner
	if got := r.RunSize(0); got != 0 {
		t.Errorf("RunSize(0) = %d, want 0", got)
	}
	r.RunSize(5)
	if got := r.RunSize(0); got != 0 {
		t.Errorf("RunSize(0) after RunSize(5) = %d, want 0", got)
	}
}
