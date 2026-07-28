package fletcher16

import "testing"

func TestChecksum(t *testing.T) {
	cases := []struct {
		in   string
		want uint16
	}{
		{"abcde", 0xC8F0},
		{"", 0x0000},
	}
	for _, c := range cases {
		if got := Checksum([]byte(c.in)); got != c.want {
			t.Errorf("Checksum(%q)=%#04x; want %#04x", c.in, got, c.want)
		}
	}
}
