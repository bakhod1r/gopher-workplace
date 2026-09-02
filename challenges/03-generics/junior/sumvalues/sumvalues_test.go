package sumvalues

import "testing"

func TestSumValues(t *testing.T) {
	if got := SumValues(map[string]int{"a": 1, "b": 2}); got != 3 {
		t.Errorf("SumValues = %v, want 3", got)
	}
	if got := SumValues(map[int]float64{1: 0.5, 2: 0.5}); got != 1 {
		t.Errorf("SumValues = %v, want 1", got)
	}
	if got := SumValues(map[string]int{}); got != 0 {
		t.Errorf("SumValues(empty) = %v, want 0", got)
	}
	if got := SumValues(map[string]int(nil)); got != 0 {
		t.Errorf("SumValues(nil) = %v, want 0", got)
	}
}
