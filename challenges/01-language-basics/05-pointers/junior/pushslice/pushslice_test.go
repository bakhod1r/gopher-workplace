package pushslice

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	var xs []int
	Push(&xs, 1)
	Push(&xs, 2)
	if !reflect.DeepEqual(xs, []int{1, 2}) {
		t.Errorf("=%v want [1 2]", xs)
	}
}
