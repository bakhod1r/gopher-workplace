package filtergen

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if got := Filter([]int{1, 2, 3, 4}, isEven); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("Filter([]int{1, 2, 3, 4}, isEven) = %v, want [2 4]", got)
	}
	if got := Filter([]int{1, 3}, isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Filter([]int{1, 3}, isEven) = %v, want []", got)
	}
	nonEmpty := func(s string) bool { return s != "" }
	if got := Filter([]string{"", "a"}, nonEmpty); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Filter([]string{\"\", \"a\"}, nonEmpty) = %v, want [a]", got)
	}
}

func TestFilterDoesNotMutate(t *testing.T) {
	in := []int{1, 2, 3}
	Filter(in, func(n int) bool { return n == 2 })
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input mutated: %v, want [1 2 3]", in)
	}
}
