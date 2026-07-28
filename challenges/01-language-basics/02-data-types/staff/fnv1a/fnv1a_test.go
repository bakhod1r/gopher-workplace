package fnv1a

import "testing"

func TestHash(t *testing.T) {
	cases := []struct {
		s    string
		want uint32
	}{
		{"", 0x811c9dc5},
		{"a", 0xe40c292c},
		{"foobar", 0xbf9cf968},
	}
	for _, c := range cases {
		if got := Hash([]byte(c.s)); got != c.want {
			t.Errorf("Hash(%q)=%#08x; want %#08x", c.s, got, c.want)
		}
	}
}
