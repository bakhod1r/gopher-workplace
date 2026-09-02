package sorteduniquegen

import (
	"reflect"
	"testing"
)

func TestSortedUnique(t *testing.T) {
	if got := SortedUnique([]int{3, 1, 3}); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("SortedUnique([]int{3, 1, 3}) = %v, want [1 3]", got)
	}
	if got := SortedUnique([]string{"b", "a", "b"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("SortedUnique = %v, want [a b]", got)
	}
	if got := SortedUnique([]int{2, 2, 2}); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("SortedUnique([]int{2, 2, 2}) = %v, want [2]", got)
	}
	if got := SortedUnique([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("SortedUnique([]int{}) = %v, want []", got)
	}
}
