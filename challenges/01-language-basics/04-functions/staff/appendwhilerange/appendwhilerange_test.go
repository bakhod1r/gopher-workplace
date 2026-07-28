package appendwhilerange

import (
	"reflect"
	"testing"
)

func TestDupAll(t *testing.T) {
	got := DupAll([]int{3, 5})
	if !reflect.DeepEqual(got, []int{3, 6, 5, 10}) {
		t.Errorf("=%v want [3 6 5 10]", got)
	}
}
