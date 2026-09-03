package bnignoredbug

import (
	"reflect"
	"testing"
)

func TestRunHonoursN(t *testing.T) {
	var seen []int
	if got := Run(3, func(i int) { seen = append(seen, i) }); got != 3 {
		t.Errorf("Run(3) = %d, want 3", got)
	}
	if !reflect.DeepEqual(seen, []int{0, 1, 2}) {
		t.Errorf("indexes = %v, want [0 1 2]", seen)
	}
}

func TestRunLargeN(t *testing.T) {
	sum := 0
	if got := Run(1000, func(i int) { sum += i }); got != 1000 {
		t.Errorf("Run(1000) = %d, want 1000", got)
	}
	if sum != 499500 {
		t.Errorf("sum = %d, want 499500", sum)
	}
}

func TestRunNonPositive(t *testing.T) {
	for _, n := range []int{0, -1} {
		calls := 0
		if got := Run(n, func(int) { calls++ }); got != 0 || calls != 0 {
			t.Errorf("Run(%d) = %d with %d calls, want 0 and 0", n, got, calls)
		}
	}
}

func TestPerOpReflectsTheRealWorkDone(t *testing.T) {
	// 300ns spread over 3 iterations is 100 ns/op. A body that runs once
	// reports 300 and claims the code is three times slower than it is.
	if got := PerOp(300, 3, func(int) {}); got != 100 {
		t.Errorf("PerOp = %d, want 100", got)
	}
}
