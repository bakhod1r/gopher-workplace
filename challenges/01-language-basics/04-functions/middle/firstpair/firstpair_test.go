package firstpair

import "testing"

func TestFindPairSum(t *testing.T) {
	i, j, ok := FindPairSum([]int{1, 2, 3, 4}, 6)
	if !ok || i != 1 || j != 3 {
		t.Errorf("=%d,%d,%v want 1,3,true", i, j, ok)
	}
	if _, _, ok := FindPairSum([]int{1, 2}, 100); ok {
		t.Errorf("no pair should be false")
	}
}
