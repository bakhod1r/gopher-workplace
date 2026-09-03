package satclampbug

import (
	"math"
	"testing"
)

func TestAddSatHigh(t *testing.T) {
	if got := AddSat(int8(100), int8(100)); got != 127 {
		t.Errorf("AddSat = %d, want 127", got)
	}
}

func TestAddSatLow(t *testing.T) {
	if got := AddSat(int8(-100), int8(-100)); got != -128 {
		t.Errorf("AddSat = %d, want -128", got)
	}
}

func TestAddSatNoClamp(t *testing.T) {
	if got := AddSat(3, 4); got != 7 {
		t.Errorf("AddSat = %d, want 7", got)
	}
	if got := AddSat(int8(-5), int8(2)); got != -3 {
		t.Errorf("AddSat = %d, want -3", got)
	}
}

func TestAddSatWideTypes(t *testing.T) {
	if got := AddSat(int32(math.MaxInt32), int32(1)); got != math.MaxInt32 {
		t.Errorf("AddSat = %d, want %d", got, int32(math.MaxInt32))
	}
	if got := AddSat(int64(math.MinInt64), int64(-1)); got != math.MinInt64 {
		t.Errorf("AddSat = %d, want %d", got, int64(math.MinInt64))
	}
}
