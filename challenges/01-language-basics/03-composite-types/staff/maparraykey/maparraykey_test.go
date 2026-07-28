package maparraykey

import "testing"

func TestCountCells(t *testing.T) {
	m := CountCells([][2]int{{1, 2}, {1, 2}, {2, 1}})
	if m[[2]int{1, 2}] != 2 {
		t.Errorf("count[{1,2}]=%d; want 2", m[[2]int{1, 2}])
	}
	if m[[2]int{2, 1}] != 1 {
		t.Errorf("count[{2,1}]=%d; want 1", m[[2]int{2, 1}])
	}
}
