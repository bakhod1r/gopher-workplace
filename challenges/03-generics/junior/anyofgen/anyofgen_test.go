package anyofgen

import "testing"

func TestAny(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if !Any([]int{1, 2}, isEven) {
		t.Error("Any([]int{1, 2}, isEven) = false, want true")
	}
	if Any([]int{1, 3}, isEven) {
		t.Error("Any([]int{1, 3}, isEven) = true, want false")
	}
	if Any([]int{}, isEven) {
		t.Error("Any([]int{}, isEven) = true, want false")
	}
}

func TestAll(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if !All([]int{2, 4}, isEven) {
		t.Error("All([]int{2, 4}, isEven) = false, want true")
	}
	if All([]int{2, 3}, isEven) {
		t.Error("All([]int{2, 3}, isEven) = true, want false")
	}
	if !All([]int{}, isEven) {
		t.Error("All([]int{}, isEven) = false, want true")
	}
}
