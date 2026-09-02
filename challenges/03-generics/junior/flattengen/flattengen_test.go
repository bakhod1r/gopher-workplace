package flattengen

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	if got := Flatten([][]int{{1, 2}, {3}}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Flatten([][]int{{1, 2}, {3}}) = %v, want [1 2 3]", got)
	}
	if got := Flatten([][]string{{}, {"a"}}); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Flatten([][]string{{}, {\"a\"}}) = %v, want [a]", got)
	}
	if got := Flatten([][]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Flatten([][]int{}) = %v, want []", got)
	}
	if got := Flatten([][]int{{}, {}}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Flatten([][]int{{}, {}}) = %v, want []", got)
	}
}
