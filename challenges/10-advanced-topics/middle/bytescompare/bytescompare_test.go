package bytescompare

import (
	"bytes"
	"testing"
)

var sink bool

func TestHasPrefix(t *testing.T) {
	cases := []struct {
		b      string
		prefix string
		want   bool
	}{
		{"hello", "he", true},
		{"hello", "hello", true},
		{"hello", "hello!", false},
		{"hello", "", true},
		{"", "x", false},
		{"hello", "ho", false},
	}
	for _, c := range cases {
		if got := HasPrefix([]byte(c.b), c.prefix); got != c.want {
			t.Errorf("HasPrefix(%q, %q) = %v, want %v", c.b, c.prefix, got, c.want)
		}
	}
}

func TestHasPrefixAllocatesNothing(t *testing.T) {
	b := bytes.Repeat([]byte("payload"), 128)
	if n := testing.AllocsPerRun(200, func() { sink = HasPrefix(b, "payload") }); n != 0 {
		t.Errorf("HasPrefix made %v allocations, want 0: do not convert b to a string", n)
	}
}
