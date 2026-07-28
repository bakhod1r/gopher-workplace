package widthmismatch

import "testing"

func TestAsU64(t *testing.T) {
	x := int64(-1) // all 64 bits set
	if got := AsU64(x); got != 0xffffffffffffffff {
		t.Errorf("=%#x want 0xffffffffffffffff (read only 32 bits?)", got)
	}
}
