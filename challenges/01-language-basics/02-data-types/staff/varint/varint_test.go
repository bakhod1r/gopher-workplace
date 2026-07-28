package varint

import "testing"

func TestDecode(t *testing.T) {
	cases := []struct {
		b   []byte
		val uint64
		n   int
	}{
		{[]byte{0x00}, 0, 1},
		{[]byte{0x7F}, 127, 1},
		{[]byte{0xAC, 0x02}, 300, 2},
		{[]byte{0x80, 0x01}, 128, 2},
		{[]byte{0xFF, 0xFF, 0x03}, 65535, 3},
		{[]byte{0x80}, 0, 0}, // truncated
	}
	for _, c := range cases {
		v, n := Decode(c.b)
		if v != c.val || n != c.n {
			t.Errorf("Decode(% x)=(%d,%d); want (%d,%d)", c.b, v, n, c.val, c.n)
		}
	}
}
