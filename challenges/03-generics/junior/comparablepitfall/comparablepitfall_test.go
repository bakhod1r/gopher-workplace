package comparablepitfall

import "testing"

func TestCountDistinct(t *testing.T) {
	if got := CountDistinct([]int{1, 1, 2}); got != 2 {
		t.Errorf("CountDistinct([]int{1, 1, 2}) = %d, want 2", got)
	}
	if got := CountDistinct([]string{"a", "b", "a"}); got != 2 {
		t.Errorf("CountDistinct = %d, want 2", got)
	}
	if got := CountDistinct([]int{}); got != 0 {
		t.Errorf("CountDistinct([]) = %d, want 0", got)
	}
}

func TestCountDistinctStructs(t *testing.T) {
	pts := []Point{{1, 2}, {1, 2}, {3, 4}}
	if got := CountDistinct(pts); got != 2 {
		t.Errorf("CountDistinct(points) = %d, want 2 (comparable structs compare field-wise)", got)
	}
}
