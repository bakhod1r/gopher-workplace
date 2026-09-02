package equalbygen

import (
	"strconv"
	"testing"
)

func TestEqualBy(t *testing.T) {
	matches := func(n int, s string) bool { return strconv.Itoa(n) == s }
	if !EqualBy([]int{1, 2}, []string{"1", "2"}, matches) {
		t.Error("EqualBy = false, want true")
	}
	if EqualBy([]int{1}, []string{"2"}, matches) {
		t.Error("EqualBy = true, want false")
	}
	if EqualBy([]int{1}, []string{"1", "2"}, matches) {
		t.Error("EqualBy with different lengths = true, want false")
	}
	if !EqualBy([]int(nil), []string{}, matches) {
		t.Error("EqualBy(nil, []) = false, want true")
	}
}

func TestEqualByTolerance(t *testing.T) {
	close := func(a, b float64) bool {
		d := a - b
		if d < 0 {
			d = -d
		}
		return d < 0.001
	}
	if !EqualBy([]float64{1.0}, []float64{1.0005}, close) {
		t.Error("EqualBy with a tolerance = false, want true")
	}
}
