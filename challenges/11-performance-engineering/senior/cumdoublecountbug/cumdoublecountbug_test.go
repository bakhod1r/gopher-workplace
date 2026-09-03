package cumdoublecountbug

import (
	"reflect"
	"testing"
)

func TestCumSumCreditsWholeStack(t *testing.T) {
	got := CumSum([]Sample{{[]string{"main", "a", "b"}, 5}})
	want := map[string]int64{"main": 5, "a": 5, "b": 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CumSum = %v, want %v", got, want)
	}
}

func TestCumSumRecursionCountedOnce(t *testing.T) {
	got := CumSum([]Sample{{[]string{"main", "rec", "rec", "rec"}, 6}})
	want := map[string]int64{"main": 6, "rec": 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CumSum = %v, want %v — recursion must not multiply the value", got, want)
	}
}

func TestCumSumDeepRecursion(t *testing.T) {
	stack := []string{"main"}
	for i := 0; i < 200; i++ {
		stack = append(stack, "descend")
	}
	got := CumSum([]Sample{{stack, 1}})
	if got["descend"] != 1 {
		t.Errorf("descend = %d, want 1 — a 200-frame recursion must still be credited once", got["descend"])
	}
}

func TestCumSumMutualRecursion(t *testing.T) {
	got := CumSum([]Sample{{[]string{"a", "b", "a", "b"}, 4}})
	want := map[string]int64{"a": 4, "b": 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CumSum = %v, want %v", got, want)
	}
}

func TestCumSumAcrossSamples(t *testing.T) {
	got := CumSum([]Sample{
		{[]string{"main", "a"}, 3},
		{[]string{"main", "b"}, 4},
	})
	want := map[string]int64{"main": 7, "a": 3, "b": 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CumSum = %v, want %v", got, want)
	}
}

func TestCumSumSkipsJunk(t *testing.T) {
	got := CumSum([]Sample{{[]string{"a"}, 0}, {nil, 5}})
	if got == nil || len(got) != 0 {
		t.Errorf("CumSum = %v, want empty non-nil map", got)
	}
}
