package ratio

import (
	"math"
	"testing"
)

func TestRatio(t *testing.T) {
	got := Value()
	want := 233.0 / 144.0
	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("Value()=%v; want %v", got, want)
	}
}
