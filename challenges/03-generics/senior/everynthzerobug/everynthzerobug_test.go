package everynthzerobug

import (
	"reflect"
	"testing"
)

func TestEveryNthStartsAtZero(t *testing.T) {
	got := EveryNth([]int{0, 1, 2, 3}, 2)
	if !reflect.DeepEqual(got, []int{0, 2}) {
		t.Errorf("EveryNth = %v, want [0 2]", got)
	}
}

func TestEveryNthStepOne(t *testing.T) {
	got := EveryNth([]int{1, 2, 3}, 1)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("EveryNth = %v, want [1 2 3]", got)
	}
}

func TestEveryNthBadStep(t *testing.T) {
	if got := EveryNth([]int{1, 2}, 0); len(got) != 0 {
		t.Errorf("EveryNth = %v, want []", got)
	}
}
