package everynthgen

import (
	"reflect"
	"testing"
)

func TestEveryNth(t *testing.T) {
	if got := EveryNth([]int{1, 2, 3, 4, 5}, 2); !reflect.DeepEqual(got, []int{1, 3, 5}) {
		t.Errorf("EveryNth(2) = %v, want [1 3 5]", got)
	}
	if got := EveryNth([]int{1, 2}, 1); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("EveryNth(1) = %v, want [1 2]", got)
	}
	if got := EveryNth([]int{1, 2, 3}, 5); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("EveryNth(5) = %v, want [1]", got)
	}
}

func TestEveryNthGuards(t *testing.T) {
	if got := EveryNth([]int{1, 2}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("EveryNth(0) = %v, want []", got)
	}
	if got := EveryNth([]int{1, 2}, -1); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("EveryNth(-1) = %v, want []", got)
	}
	if got := EveryNth([]int{}, 2); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("EveryNth(empty) = %v, want []", got)
	}
}
