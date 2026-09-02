package interleavegen

import (
	"reflect"
	"testing"
)

func TestInterleave(t *testing.T) {
	if got := Interleave([]int{1, 3}, []int{2, 4}); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("Interleave = %v, want [1 2 3 4]", got)
	}
	if got := Interleave([]int{1, 2, 3}, []int{9}); !reflect.DeepEqual(got, []int{1, 9, 2, 3}) {
		t.Errorf("Interleave = %v, want [1 9 2 3]", got)
	}
	if got := Interleave([]int{1}, []int{8, 9}); !reflect.DeepEqual(got, []int{1, 8, 9}) {
		t.Errorf("Interleave = %v, want [1 8 9]", got)
	}
	if got := Interleave([]int{}, []int{1}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Interleave = %v, want [1]", got)
	}
	if got := Interleave([]int{}, []int{}); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Interleave = %v, want []", got)
	}
}
