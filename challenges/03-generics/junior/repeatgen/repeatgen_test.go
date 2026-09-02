package repeatgen

import (
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	if got := Repeat([]int{1, 2}, 2); !reflect.DeepEqual(got, []int{1, 2, 1, 2}) {
		t.Errorf("Repeat([]int{1, 2}, 2) = %v, want [1 2 1 2]", got)
	}
	if got := Repeat([]string{"a"}, 3); !reflect.DeepEqual(got, []string{"a", "a", "a"}) {
		t.Errorf("Repeat([]string{\"a\"}, 3) = %v, want [a a a]", got)
	}
	if got := Repeat([]int{1}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Repeat([]int{1}, 0) = %v, want []", got)
	}
	if got := Repeat([]int{1}, -1); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Repeat([]int{1}, -1) = %v, want []", got)
	}
}
