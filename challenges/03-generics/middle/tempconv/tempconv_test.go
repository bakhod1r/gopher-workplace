package tempconv

import "testing"

type Celsius float64

func TestToFloat(t *testing.T) {
	if got := ToFloat(Celsius(20)); got != 20.0 {
		t.Errorf("ToFloat = %v, want 20", got)
	}
	if got := ToFloat(2.5); got != 2.5 {
		t.Errorf("ToFloat = %v, want 2.5", got)
	}
}

func TestFromFloat(t *testing.T) {
	if got := FromFloat[Celsius](20); got != Celsius(20) {
		t.Errorf("FromFloat = %v, want 20", got)
	}
	var c Celsius = FromFloat[Celsius](-5)
	if c != Celsius(-5) {
		t.Errorf("FromFloat = %v, want -5", c)
	}
}

func TestRescale(t *testing.T) {
	double := func(f float64) float64 { return f * 2 }
	got := Rescale(Celsius(20), double)
	if got != Celsius(40) {
		t.Errorf("Rescale = %v, want 40", got)
	}
	var c Celsius = got
	if c != Celsius(40) {
		t.Errorf("Rescale lost the named type: %v", c)
	}
}
