package floatbits

import (
	"math"
	"testing"
)

func TestBits(t *testing.T) {
	for _, f := range []float64{1.5, 0.1, -2.0, 3.14159} {
		if got := Bits(f); got != math.Float64bits(f) {
			t.Errorf("Bits(%v)=%#x want %#x", f, got, math.Float64bits(f))
		}
	}
}
