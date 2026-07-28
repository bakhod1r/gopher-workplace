package stackpush

import (
	"reflect"
	"testing"
)

func TestReverseInts(t *testing.T) {
	if got := ReverseInts([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("=%v want [3 2 1]", got)
	}
}
