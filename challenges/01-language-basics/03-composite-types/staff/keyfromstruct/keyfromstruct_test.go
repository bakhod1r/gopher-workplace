package keyfromstruct

import "testing"

func TestCount(t *testing.T) {
	m := Count([]Point{{1, 2}, {1, 2}, {2, 1}})
	if m[Point{1, 2}] != 2 {
		t.Errorf("count{1,2}=%d; want 2", m[Point{1, 2}])
	}
	if m[Point{2, 1}] != 1 {
		t.Errorf("count{2,1}=%d; want 1", m[Point{2, 1}])
	}
}
