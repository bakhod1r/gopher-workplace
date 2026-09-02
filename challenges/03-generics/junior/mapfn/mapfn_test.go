package mapfn

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMapSameType(t *testing.T) {
	double := func(n int) int { return n * 2 }
	if got := Map([]int{1, 2}, double); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("Map([]int{1, 2}, double) = %v, want [2 4]", got)
	}
	if got := Map([]int{}, double); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Map([]int{}, double) = %v, want []", got)
	}
}

func TestMapChangesType(t *testing.T) {
	if got := Map([]int{1, 2}, strconv.Itoa); !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("Map([]int{1, 2}, strconv.Itoa) = %v, want [1 2]", got)
	}
}
