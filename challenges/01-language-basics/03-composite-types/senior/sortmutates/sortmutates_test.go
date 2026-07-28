package sortmutates

import (
	"reflect"
	"testing"
)

func TestSortedCopy(t *testing.T) {
	xs := []int{3, 1, 2}
	got := SortedCopy(xs)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("SortedCopy=%v", got)
	}
	if !reflect.DeepEqual(xs, []int{3, 1, 2}) {
		t.Errorf("input mutated: %v", xs)
	}
}
