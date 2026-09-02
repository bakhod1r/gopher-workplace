package takewhilegen

import (
	"reflect"
	"testing"
)

func TestTakeWhile(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if got := TakeWhile([]int{2, 4, 5, 6}, isEven); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("TakeWhile = %v, want [2 4] (stop at the first failure)", got)
	}
	if got := TakeWhile([]int{1, 2}, isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("TakeWhile = %v, want []", got)
	}
	if got := TakeWhile([]int{2, 4}, isEven); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("TakeWhile = %v, want [2 4]", got)
	}
	if got := TakeWhile([]int{}, isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("TakeWhile(empty) = %v, want []", got)
	}
}
