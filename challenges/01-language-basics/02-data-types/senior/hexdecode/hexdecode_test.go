package hexdecode

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		s    string
		want []byte
		ok   bool
	}{
		{"ff", []byte{0xff}, true},
		{"1a2b", []byte{0x1a, 0x2b}, true},
		{"476F", []byte("Go"), true},
		{"abc", nil, false},
		{"xy", nil, false},
	}
	for _, c := range cases {
		got, ok := Decode(c.s)
		if ok != c.ok || (ok && !bytes.Equal(got, c.want)) {
			t.Errorf("Decode(%q)=(%v,%v); want (%v,%v)", c.s, got, ok, c.want, c.ok)
		}
	}
}
