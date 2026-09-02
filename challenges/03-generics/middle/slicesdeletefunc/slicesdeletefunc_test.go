package slicesdeletefunc

import (
	"reflect"
	"testing"
)

func TestPurge(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if got := Purge([]int{1, 2, 3, 4}, isEven); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Purge = %v, want [1 3]", got)
	}
	if got := Purge([]int{2}, isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Purge = %v, want []", got)
	}
	if got := Purge([]int(nil), isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Purge(nil) = %v, want []", got)
	}
}

func TestPurgeDoesNotMutate(t *testing.T) {
	in := []int{1, 2, 3}
	Purge(in, func(n int) bool { return n == 2 })
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input mutated: %v, want [1 2 3]", in)
	}
}
