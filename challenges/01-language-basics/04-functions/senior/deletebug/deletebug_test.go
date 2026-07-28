package deletebug

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	got := RemoveAt([]int{10, 20, 30, 40}, 1)
	if !reflect.DeepEqual(got, []int{10, 30, 40}) {
		t.Errorf("=%v want [10 30 40]", got)
	}
}
