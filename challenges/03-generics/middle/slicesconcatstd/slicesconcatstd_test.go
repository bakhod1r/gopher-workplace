package slicesconcatstd

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	if got := Join([]int{1}, []int{2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Join = %v, want [1 2 3]", got)
	}
	if got := Join([]string{"a"}); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Join = %v, want [a]", got)
	}
	if got := Join[int](); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Join() = %v, want []", got)
	}
}

func TestJoinDoesNotAlias(t *testing.T) {
	a := []int{1}
	got := Join(a, []int{2})
	got[0] = 99
	if a[0] != 1 {
		t.Errorf("result aliases the input: a = %v", a)
	}
}
