package intersectgen

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	if got := Intersect([]int{1, 2, 2, 3}, []int{2, 3}); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("Intersect = %v, want [2 3]", got)
	}
	if got := Intersect([]int{1}, []int{2}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Intersect = %v, want []", got)
	}
	if got := Intersect([]int{}, []int{1}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Intersect = %v, want []", got)
	}
	if got := Intersect([]int{3, 1}, []int{1, 3}); !reflect.DeepEqual(got, []int{3, 1}) {
		t.Errorf("Intersect = %v, want [3 1] (order comes from a)", got)
	}
}

func TestIntersectStrings(t *testing.T) {
	got := Intersect([]string{"a", "b", "a"}, []string{"a"})
	if !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Intersect = %v, want [a]", got)
	}
}
