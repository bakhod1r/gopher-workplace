package appendall

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	base := []int{1, 2}
	got := Concat(base, 3, 4, 5)
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("=%v want [1 2 3 4 5]", got)
	}
	if !reflect.DeepEqual(base, []int{1, 2}) {
		t.Errorf("base mutated: %v", base)
	}
}
