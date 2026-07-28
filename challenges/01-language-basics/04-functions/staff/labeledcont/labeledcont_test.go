package labeledcont

import "testing"

func TestCleanRows(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{1, -1, 2},
		{4, 5},
	}
	if got := CleanRows(grid); got != 2 {
		t.Errorf("=%d want 2 (rows 0 and 2)", got)
	}
}
