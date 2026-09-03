package stringbytesconv

import (
	"strings"
	"testing"
)

var sink int
var sinkB bool

func TestCountByte(t *testing.T) {
	cases := []struct {
		s    string
		b    byte
		want int
	}{
		{"hello", 'l', 2},
		{"hello", 'z', 0},
		{"", 'a', 0},
		{"aaa", 'a', 3},
	}
	for _, c := range cases {
		if got := CountByte(c.s, c.b); got != c.want {
			t.Errorf("CountByte(%q, %q) = %d, want %d", c.s, c.b, got, c.want)
		}
	}
}

func TestCountByteCountsBytesNotRunes(t *testing.T) {
	// "é" is two bytes: 0xC3 0xA9.
	if got := CountByte("é", 0xC3); got != 1 {
		t.Errorf("CountByte = %d, want 1", got)
	}
}

func TestCountByteDoesNotAllocate(t *testing.T) {
	s := strings.Repeat("abcdefgh", 500)
	allocs := testing.AllocsPerRun(50, func() { sink = CountByte(s, 'c') })
	if allocs != 0 {
		t.Errorf("CountByte made %v allocations, want 0 — a []byte conversion copies the string", allocs)
	}
}

func TestHasPrefixByte(t *testing.T) {
	if !HasPrefixByte("go", 'g') {
		t.Error("HasPrefixByte(\"go\", 'g') = false, want true")
	}
	if HasPrefixByte("go", 'o') {
		t.Error("HasPrefixByte(\"go\", 'o') = true, want false")
	}
	if HasPrefixByte("", 'a') {
		t.Error("HasPrefixByte(\"\", 'a') = true, want false")
	}
}

func TestHasPrefixByteDoesNotAllocate(t *testing.T) {
	s := strings.Repeat("abcdefgh", 500)
	allocs := testing.AllocsPerRun(50, func() { sinkB = HasPrefixByte(s, 'a') })
	if allocs != 0 {
		t.Errorf("HasPrefixByte made %v allocations, want 0", allocs)
	}
}
