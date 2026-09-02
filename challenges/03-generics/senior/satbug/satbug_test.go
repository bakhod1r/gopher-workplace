package satbug

import "testing"

func TestSatAddSaturates(t *testing.T) {
	const maxUint64 = ^uint64(0)
	if got := SatAdd(maxUint64, 1); got != maxUint64 {
		t.Errorf("SatAdd(max, 1) = %v, want %v (it wrapped)", got, maxUint64)
	}
	if got := SatAdd(maxUint64, maxUint64); got != maxUint64 {
		t.Errorf("SatAdd(max, max) = %v, want %v", got, maxUint64)
	}
	maxUint := ^uint(0)
	if got := SatAdd(maxUint, uint(5)); got != maxUint {
		t.Errorf("SatAdd(maxUint, 5) = %v, want %v", got, maxUint)
	}
}

func TestSatAddNormal(t *testing.T) {
	if got := SatAdd(uint(1), uint(2)); got != 3 {
		t.Errorf("SatAdd(1, 2) = %v, want 3", got)
	}
	if got := SatAdd(uint(0), uint(0)); got != 0 {
		t.Errorf("SatAdd(0, 0) = %v, want 0", got)
	}
	const maxUint64 = ^uint64(0)
	if got := SatAdd(maxUint64-2, 2); got != maxUint64 {
		t.Errorf("SatAdd(max-2, 2) = %v, want %v (exactly max, no overflow)", got, maxUint64)
	}
}
