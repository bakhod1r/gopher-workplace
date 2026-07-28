package sums

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	got := SumTenths(10)
	if math.Abs(got-1.0) > 1e-9 {
		t.Fatalf("SumTenths(10)=%v; want ~1.0 within 1e-9", got)
	}
	got = SumTenths(100)
	if math.Abs(got-10.0) > 1e-9 {
		t.Fatalf("SumTenths(100)=%v; want ~10.0 within 1e-9", got)
	}
}
