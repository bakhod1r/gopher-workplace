package slicesequalstd

import "testing"

func TestSameOrder(t *testing.T) {
	if !SameOrder([]int{1, 2}, []int{1, 2}) {
		t.Error("SameOrder([1 2], [1 2]) = false, want true")
	}
	if SameOrder([]int{1, 2}, []int{2, 1}) {
		t.Error("SameOrder([1 2], [2 1]) = true, want false")
	}
	if SameOrder([]int{1}, []int{1, 2}) {
		t.Error("SameOrder([1], [1 2]) = true, want false")
	}
	if !SameOrder([]int(nil), []int{}) {
		t.Error("SameOrder(nil, []) = false, want true")
	}
	if !SameOrder([]string{"a"}, []string{"a"}) {
		t.Error(`SameOrder([a], [a]) = false, want true`)
	}
}
