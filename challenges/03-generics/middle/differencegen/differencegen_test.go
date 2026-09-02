package differencegen

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	if got := Difference([]int{1, 2, 2, 3}, []int{2}); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Difference = %v, want [1 3]", got)
	}
	if got := Difference([]int{1}, []int{1}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Difference = %v, want []", got)
	}
	if got := Difference([]int{}, []int{1}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Difference = %v, want []", got)
	}
	if got := Difference([]int{3, 1}, []int{}); !reflect.DeepEqual(got, []int{3, 1}) {
		t.Errorf("Difference = %v, want [3 1] (first-seen order)", got)
	}
	if got := Difference([]string{"a", "a"}, []string{"b"}); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Difference = %v, want [a] (distinct)", got)
	}
}

func TestDifferenceDoesNotMutate(t *testing.T) {
	a := []int{1, 2}
	b := []int{2}
	Difference(a, b)
	if !reflect.DeepEqual(a, []int{1, 2}) || !reflect.DeepEqual(b, []int{2}) {
		t.Errorf("inputs mutated: a=%v b=%v", a, b)
	}
}
