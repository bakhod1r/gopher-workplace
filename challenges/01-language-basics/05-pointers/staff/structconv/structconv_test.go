package structconv

import "testing"

func TestToVec(t *testing.T) {
	v := ToVec(&Point{X: 3, Y: 4})
	if v.X != 3 || v.Y != 4 {
		t.Errorf("=%+v want {3 4}", v)
	}
}
