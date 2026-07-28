package deferloop

import (
	"reflect"
	"testing"
)

func TestCloseOrder(t *testing.T) {
	if got := CloseOrder(3); !reflect.DeepEqual(got, []int{2, 1, 0}) {
		t.Errorf("=%v want [2 1 0]", got)
	}
}
