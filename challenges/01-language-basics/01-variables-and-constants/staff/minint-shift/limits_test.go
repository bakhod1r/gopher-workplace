package limits

import (
	"math"
	"testing"
)

func TestMinInt(t *testing.T) {
	if int64(MinInt) != math.MinInt64 {
		t.Fatalf("MinInt=%d; want %d", int64(MinInt), int64(math.MinInt64))
	}
}

func TestAsymmetry(t *testing.T) {
	if !SymmetricTo() {
		t.Error("MinInt must have no positive counterpart")
	}
}
