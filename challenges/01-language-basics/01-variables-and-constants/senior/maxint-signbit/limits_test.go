package limits

import (
	"math"
	"testing"
)

func TestMaxInt(t *testing.T) {
	if MaxInt <= 0 {
		t.Fatalf("MaxInt=%d must be positive", MaxInt)
	}
	if int64(MaxInt) != math.MaxInt64 && int64(MaxInt) != math.MaxInt32 {
		t.Fatalf("MaxInt=%d not a machine max", MaxInt)
	}
}

func TestOverflows(t *testing.T) {
	if !Overflows() {
		t.Error("MaxInt+1 must wrap negative")
	}
}
