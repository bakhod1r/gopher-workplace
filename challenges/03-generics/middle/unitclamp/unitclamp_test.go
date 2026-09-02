package unitclamp

import "testing"

type Meters float64
type Seconds float64

func TestClampUnit(t *testing.T) {
	if got := ClampUnit(Meters(5), Meters(0), Meters(3)); got != Meters(3) {
		t.Errorf("ClampUnit = %v, want 3", got)
	}
	if got := ClampUnit(Seconds(-1), Seconds(0), Seconds(3)); got != Seconds(0) {
		t.Errorf("ClampUnit = %v, want 0", got)
	}
	if got := ClampUnit(Meters(2), Meters(0), Meters(3)); got != Meters(2) {
		t.Errorf("ClampUnit = %v, want 2", got)
	}
	if got := ClampUnit(2.0, 0.0, 3.0); got != 2.0 {
		t.Errorf("ClampUnit(plain float) = %v, want 2", got)
	}
}

func TestClampUnitKeepsNamedType(t *testing.T) {
	var m Meters = ClampUnit(Meters(9), Meters(0), Meters(1))
	if m != Meters(1) {
		t.Errorf("result = %v, want Meters(1)", m)
	}
}
