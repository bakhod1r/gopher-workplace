package uniquegen

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	if got := Unique([]int{1, 2, 1}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Unique([]int{1, 2, 1}) = %v, want [1 2]", got)
	}
	if got := Unique([]int{3, 1, 3, 2, 1}); !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Errorf("Unique([]int{3, 1, 3, 2, 1}) = %v, want [3 1 2]", got)
	}
	if got := Unique([]string{"a", "a", "b"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Unique([]string{\"a\", \"a\", \"b\"}) = %v, want [a b]", got)
	}
	if got := Unique([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Unique([]int{}) = %v, want []", got)
	}
}
