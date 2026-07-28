package color

import "testing"

func TestPack(t *testing.T) {
	cases := []struct {
		r, g, b uint8
		want    uint32
	}{
		{0xFF, 0x00, 0x00, 0xFF0000},
		{0x00, 0xFF, 0x00, 0x00FF00},
		{0x12, 0x34, 0x56, 0x123456},
	}
	for _, c := range cases {
		if got := Pack(c.r, c.g, c.b); got != c.want {
			t.Errorf("Pack(%#x,%#x,%#x)=%#06x; want %#06x", c.r, c.g, c.b, got, c.want)
		}
	}
}

func TestRed(t *testing.T) {
	if Red(Pack(0xAB, 0, 0)) != 0xAB {
		t.Error("red roundtrip failed")
	}
}
