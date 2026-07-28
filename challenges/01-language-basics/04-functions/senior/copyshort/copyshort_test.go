package copyshort

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	xs := []int{1, 2, 3}
	got := Clone(xs)
	if !reflect.DeepEqual(got, xs) {
		t.Errorf("=%v want %v", got, xs)
	}
	got[0] = 99
	if xs[0] == 99 {
		t.Errorf("clone shares backing array")
	}
}
