package concatgen

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	if got := Concat([]int{1}, []int{2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Concat([]int{1}, []int{2, 3}) = %v, want [1 2 3]", got)
	}
	if got := Concat([]string{"a"}); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Concat([]string{\"a\"}) = %v, want [a]", got)
	}
	if got := Concat[int](); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Concat[int]() = %v, want []", got)
	}
}

func TestConcatDoesNotAliasInputs(t *testing.T) {
	a := make([]int, 1, 4)
	a[0] = 1
	got := Concat(a, []int{2})
	got[0] = 99
	if a[0] != 1 {
		t.Errorf("writing to the result changed the input: a[0] = %v, want 1", a[0])
	}
	if len(got) != 2 {
		t.Errorf("len(result) = %d, want 2", len(got))
	}
}
