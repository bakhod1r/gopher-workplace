package hexencode

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct {
		b    []byte
		want string
	}{
		{[]byte{0x00}, "00"},
		{[]byte{0xFF}, "ff"},
		{[]byte{0x1a, 0x2b}, "1a2b"},
		{[]byte("Go"), "476f"},
		{nil, ""},
	}
	for _, c := range cases {
		if got := Encode(c.b); got != c.want {
			t.Errorf("Encode(%v)=%q; want %q", c.b, got, c.want)
		}
	}
}
