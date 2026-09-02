package tofloats

import (
	"reflect"
	"testing"
)

func TestToFloats(t *testing.T) {
	if got := ToFloats([]int{1, 2}); !reflect.DeepEqual(got, []float64{1, 2}) {
		t.Errorf("ToFloats([]int{1, 2}) = %v, want [1 2]", got)
	}
	if got := ToFloats([]int64{7}); !reflect.DeepEqual(got, []float64{7}) {
		t.Errorf("ToFloats([]int64{7}) = %v, want [7]", got)
	}
	if got := ToFloats([]int{}); !reflect.DeepEqual(got, []float64{}) {
		t.Errorf("ToFloats([]int{}) = %v, want []", got)
	}
}
