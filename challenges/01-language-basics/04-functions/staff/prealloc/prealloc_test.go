package prealloc

import (
	"reflect"
	"testing"
)

func TestDoubles(t *testing.T) {
	if got := Doubles(3); !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("=%v want [2 4 6]", got)
	}
	if got := Doubles(0); len(got) != 0 {
		t.Errorf("=%v want empty", got)
	}
}
