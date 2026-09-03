package filterinplacebug

import (
	"reflect"
	"testing"
)

func even(n int) bool { return n%2 == 0 }

func TestFilterInPlaceShrinks(t *testing.T) {
	got := FilterInPlace([]int{1, 2, 3, 4}, even)
	if !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("FilterInPlace = %v, want [2 4]", got)
	}
}

func TestFilterInPlaceNoneKept(t *testing.T) {
	if got := FilterInPlace([]int{1, 3}, even); len(got) != 0 {
		t.Errorf("FilterInPlace = %v, want []", got)
	}
}

func TestFilterInPlaceAllKept(t *testing.T) {
	got := FilterInPlace([]int{2}, even)
	if !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("FilterInPlace = %v, want [2]", got)
	}
}
