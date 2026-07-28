package extract

import "testing"

func TestByteAt(t *testing.T) {
	v := uint64(0x1122334455667788)
	want := []uint8{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}
	for n, w := range want {
		if got := ByteAt(v, uint(n)); got != w {
			t.Errorf("ByteAt(v,%d)=%#x; want %#x", n, got, w)
		}
	}
}
