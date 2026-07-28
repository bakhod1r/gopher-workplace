package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	i, j, ok := TwoSum([]int{2, 7, 11, 15}, 9)
	if !ok || i != 0 || j != 1 {
		t.Errorf("got (%d,%d,%v); want (0,1,true)", i, j, ok)
	}
	i, j, ok = TwoSum([]int{3, 2, 4}, 6)
	if !ok || i != 1 || j != 2 {
		t.Errorf("got (%d,%d,%v); want (1,2,true)", i, j, ok)
	}
	if _, _, ok := TwoSum([]int{1, 2}, 100); ok {
		t.Error("no pair should be false")
	}
}
