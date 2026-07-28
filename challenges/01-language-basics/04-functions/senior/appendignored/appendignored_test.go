package appendignored

import (
	"reflect"
	"testing"
)

func TestSquares(t *testing.T) {
	got := Squares(4)
	if !reflect.DeepEqual(got, []int{1, 4, 9, 16}) {
		t.Errorf("=%v want [1 4 9 16]", got)
	}
}
