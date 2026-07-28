package bytesum

import "testing"

func TestSum(t *testing.T) {
	bs := []byte{200, 100, 50} // 350, overflows uint8
	if got := Sum(bs); got != 350 {
		t.Errorf("=%d want 350 (uint8 wrapped?)", got)
	}
}
