package bytescontains

import (
	"bytes"
	"testing"
)

var sink bool

func TestContains(t *testing.T) {
	cases := []struct {
		hay, needle string
		want        bool
	}{
		{"hello", "ell", true},
		{"hello", "hello", true},
		{"hello", "hellos", false},
		{"hello", "", true},
		{"", "", true},
		{"", "x", false},
		{"hello", "lo", true},
		{"hello", "he", true},
		{"aaa", "aab", false},
		{"abcabd", "abd", true},
	}
	for _, c := range cases {
		if got := Contains([]byte(c.hay), c.needle); got != c.want {
			t.Errorf("Contains(%q, %q) = %v, want %v", c.hay, c.needle, got, c.want)
		}
	}
}

func TestContainsMatchesStdlib(t *testing.T) {
	hay := []byte("the quick brown fox jumps over the lazy dog")
	for _, needle := range []string{"quick", "dog", "cat", "the", "g", "", "o d"} {
		want := bytes.Contains(hay, []byte(needle))
		if got := Contains(hay, needle); got != want {
			t.Errorf("Contains(_, %q) = %v, want %v", needle, got, want)
		}
	}
}

func TestContainsAllocatesNothing(t *testing.T) {
	hay := bytes.Repeat([]byte("payload-"), 256)
	if n := testing.AllocsPerRun(100, func() { sink = Contains(hay, "payload-payload") }); n != 0 {
		t.Errorf("Contains made %v allocations, want 0", n)
	}
}
