package mergesortedgen

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	if got := Merge([]int{1, 3}, []int{2, 4}); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("Merge = %v, want [1 2 3 4]", got)
	}
	if got := Merge([]int{1, 2}, []int{3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Merge = %v, want [1 2 3]", got)
	}
	if got := Merge([]int{}, []int{1}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Merge = %v, want [1]", got)
	}
	if got := Merge([]int{}, []int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Merge = %v, want []", got)
	}
	if got := Merge([]string{"a", "c"}, []string{"b"}); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Merge = %v, want [a b c]", got)
	}
}

func TestMergeKeepsDuplicates(t *testing.T) {
	if got := Merge([]int{1, 1}, []int{1}); !reflect.DeepEqual(got, []int{1, 1, 1}) {
		t.Errorf("Merge = %v, want [1 1 1]", got)
	}
}
