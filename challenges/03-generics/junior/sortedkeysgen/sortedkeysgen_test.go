package sortedkeysgen

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	if got := SortedKeys(map[string]int{"b": 1, "a": 2}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("SortedKeys = %v, want [a b]", got)
	}
	if got := SortedKeys(map[int]bool{3: true, 1: false, 2: true}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("SortedKeys = %v, want [1 2 3]", got)
	}
	if got := SortedKeys(map[string]int{}); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("SortedKeys(empty) = %v, want []", got)
	}
}
