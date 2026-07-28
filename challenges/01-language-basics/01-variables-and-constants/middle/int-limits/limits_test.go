package limits

import (
	"math"
	"testing"
)

func TestBounds(t *testing.T) {
	if MaxUint != math.MaxUint64 && MaxUint != math.MaxUint32 {
		t.Fatalf("MaxUint=%d unexpected", MaxUint)
	}
	if uint64(MaxInt) != math.MaxInt64 && int64(MaxInt) != math.MaxInt32 {
		t.Fatalf("MaxInt=%d unexpected", MaxInt)
	}
	if MinInt != -MaxInt-1 {
		t.Fatalf("MinInt=%d; want %d", MinInt, -MaxInt-1)
	}
}

func TestFits(t *testing.T) {
	if !FitsInInt(0) || !FitsInInt(uint(MaxInt)) {
		t.Error("small values should fit")
	}
	if FitsInInt(MaxUint) {
		t.Error("MaxUint must not fit in int")
	}
}
