package valuesgen

import (
	"reflect"
	"sort"
	"testing"
)

func TestValues(t *testing.T) {
	got := Values(map[string]int{"a": 1, "b": 2})
	sort.Ints(got)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Values = %v, want [1 2]", got)
	}

	dup := Values(map[string]int{"a": 1, "b": 1})
	if len(dup) != 2 {
		t.Errorf("Values kept %d values, want 2 (duplicates are not removed)", len(dup))
	}

	strs := Values(map[int]string{1: "x"})
	if !reflect.DeepEqual(strs, []string{"x"}) {
		t.Errorf("Values = %v, want [x]", strs)
	}

	if e := Values(map[string]int{}); !reflect.DeepEqual(e, []int{}) {
		t.Errorf("Values(empty) = %v, want []", e)
	}
}
