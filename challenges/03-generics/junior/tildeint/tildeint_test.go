package tildeint

import "testing"

func TestSumTempsNamedType(t *testing.T) {
	got := SumTemps([]Celsius{1, 2})
	if got != Celsius(3) {
		t.Errorf("SumTemps([]Celsius{1, 2}) = %v, want 3", got)
	}
	var want Celsius = 3
	if got != want {
		t.Errorf("result type is not Celsius: %T", got)
	}
}

func TestSumTempsPlainInt(t *testing.T) {
	if got := SumTemps([]int{5}); got != 5 {
		t.Errorf("SumTemps([]int{5}) = %v, want 5", got)
	}
	if got := SumTemps([]Celsius{}); got != 0 {
		t.Errorf("SumTemps([]Celsius{}) = %v, want 0", got)
	}
}
