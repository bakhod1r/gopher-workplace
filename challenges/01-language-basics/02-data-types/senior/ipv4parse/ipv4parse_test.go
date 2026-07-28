package ipv4parse

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		s    string
		want [4]byte
		ok   bool
	}{
		{"192.168.1.1", [4]byte{192, 168, 1, 1}, true},
		{"0.0.0.0", [4]byte{}, true},
		{"255.255.255.255", [4]byte{255, 255, 255, 255}, true},
		{"256.1.1.1", [4]byte{}, false}, // octet out of range
		{"1.2.3", [4]byte{}, false},
		{"1.2.3.4.5", [4]byte{}, false},
	}
	for _, c := range cases {
		got, ok := Parse(c.s)
		if ok != c.ok || (ok && got != c.want) {
			t.Errorf("Parse(%q)=(%v,%v); want (%v,%v)", c.s, got, ok, c.want, c.ok)
		}
	}
}
