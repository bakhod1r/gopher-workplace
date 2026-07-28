package poolreset

import "testing"

func TestReset(t *testing.T) {
	b := &Buf{Data: []byte("stale"), Len: 5}
	Reset(b)
	if b.Len != 0 {
		t.Errorf("Len=%d want 0", b.Len)
	}
	if len(b.Data) != 0 {
		t.Errorf("Data len=%d want 0 (stale contents retained)", len(b.Data))
	}
}
