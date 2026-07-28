package reduceints

import "testing"

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(a, x int) int { return a + x })
	if sum != 10 {
		t.Errorf("sum=%d want 10", sum)
	}
	prod := Reduce([]int{1, 2, 3, 4}, 1, func(a, x int) int { return a * x })
	if prod != 24 {
		t.Errorf("prod=%d want 24", prod)
	}
}
