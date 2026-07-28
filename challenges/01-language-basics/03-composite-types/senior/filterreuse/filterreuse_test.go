package filterreuse

import (
	"reflect"
	"testing"
)

func TestEvens(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6}
	got := Evens(xs)
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("Evens=%v; want [2 4 6]", got)
	}
	if !reflect.DeepEqual(xs, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("input corrupted: %v", xs)
	}
}
