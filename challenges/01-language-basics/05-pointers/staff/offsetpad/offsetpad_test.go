package offsetpad

import "testing"

func TestReadN(t *testing.T) {
	r := &Rec{Flag: true, N: 123456789}
	if got := ReadN(r); got != 123456789 {
		t.Errorf("=%d want 123456789 (ignored alignment padding?)", got)
	}
}
