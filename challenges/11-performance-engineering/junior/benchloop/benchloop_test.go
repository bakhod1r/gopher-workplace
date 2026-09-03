package benchloop

import (
	"reflect"
	"testing"
)

func TestRunCallsWorkNTimes(t *testing.T) {
	var seen []int
	if got := Run(3, func(i int) { seen = append(seen, i) }); got != 3 {
		t.Errorf("Run = %d, want 3", got)
	}
	if !reflect.DeepEqual(seen, []int{0, 1, 2}) {
		t.Errorf("indexes = %v, want [0 1 2]", seen)
	}
}

func TestRunZeroAndNegative(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		calls := 0
		if got := Run(n, func(int) { calls++ }); got != 0 || calls != 0 {
			t.Errorf("Run(%d) = %d with %d calls, want 0 and 0", n, got, calls)
		}
	}
}

func TestRunLargeN(t *testing.T) {
	sum := 0
	Run(1000, func(i int) { sum += i })
	if sum != 499500 {
		t.Errorf("sum = %d, want 499500", sum)
	}
}
