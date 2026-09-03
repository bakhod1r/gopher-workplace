package windowmaxexpirebug

import (
	"reflect"
	"testing"
)

func TestWindowMaxExpires(t *testing.T) {
	got := WindowMax([]int{5, 1, 1, 1}, 2)
	if !reflect.DeepEqual(got, []int{5, 1, 1}) {
		t.Errorf("WindowMax = %v, want [5 1 1]", got)
	}
}

func TestWindowMaxSimple(t *testing.T) {
	got := WindowMax([]int{1, 3, 2}, 2)
	if !reflect.DeepEqual(got, []int{3, 3}) {
		t.Errorf("WindowMax = %v, want [3 3]", got)
	}
}

func TestWindowMaxBadWidth(t *testing.T) {
	if got := WindowMax([]int{1}, 5); len(got) != 0 {
		t.Errorf("WindowMax = %v, want []", got)
	}
	if got := WindowMax([]int{1}, 0); len(got) != 0 {
		t.Errorf("WindowMax = %v, want []", got)
	}
}
