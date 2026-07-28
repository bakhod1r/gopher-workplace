package deferorder

import (
	"reflect"
	"testing"
)

func TestOrder(t *testing.T) {
	if got := Order(); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("=%v want [3 2 1]", got)
	}
}
