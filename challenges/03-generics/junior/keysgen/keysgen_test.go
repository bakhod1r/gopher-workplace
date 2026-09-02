package keysgen

import (
	"reflect"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	got := Keys(map[string]int{"b": 2, "a": 1})
	sort.Strings(got)
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Keys = %v, want [a b]", got)
	}

	ints := Keys(map[int]bool{1: true, 2: false})
	sort.Ints(ints)
	if !reflect.DeepEqual(ints, []int{1, 2}) {
		t.Errorf("Keys = %v, want [1 2]", ints)
	}

	if e := Keys(map[string]int{}); !reflect.DeepEqual(e, []string{}) {
		t.Errorf("Keys(empty) = %v, want []", e)
	}
}
