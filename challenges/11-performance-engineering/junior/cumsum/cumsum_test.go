package cumsum

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

func TestCumSumRecursionCountedOnce(t *testing.T) {
	got := CumSum([]Sample{{[]string{"main", "rec", "rec", "rec"}, 6}})
	want := map[string]int64{"main": 6, "rec": 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CumSum = %v, want %v (recursion must not multiply the value)", got, want)
	}
}

func TestCumSumSkipsJunk(t *testing.T) {
	got := CumSum([]Sample{
		{[]string{"a"}, 0},
		{[]string{"b"}, -1},
		{nil, 9},
	})
	if got == nil || len(got) != 0 {
		t.Errorf("CumSum = %v, want empty non-nil map", got)
	}
}
