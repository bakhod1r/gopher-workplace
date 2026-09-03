package bytesequal

import (
	"bytes"
	"testing"
)

var sink bool

func TestEqualString(t *testing.T) {
	cases := []struct {
		b, s string
		want bool
	}{
		{"hi", "hi", true},
		{"hi", "ho", false},
		{"hi", "his", false},
		{"his", "hi", false},
		{"", "", true},
		{"", "x", false},
		{"x", "", false},
	}
	for _, c := range cases {
		if got := EqualString([]byte(c.b), c.s); got != c.want {
			t.Errorf("EqualString(%q, %q) = %v, want %v", c.b, c.s, got, c.want)
		}
	}
}

func TestEqualStringNil(t *testing.T) {
	if !EqualString(nil, "") {
		t.Error("EqualString(nil, \"\") = false, want true")
	}
	if EqualString(nil, "x") {
		t.Error("EqualString(nil, \"x\") = true, want false")
	}
}

func TestEqualStringAllocatesNothing(t *testing.T) {
	b := bytes.Repeat([]byte("payload"), 512)
	s := string(b)
	if n := testing.AllocsPerRun(200, func() { sink = EqualString(b, s) }); n != 0 {
		t.Errorf("EqualString made %v allocations, want 0", n)
	}
}

func TestEqualStringDoesNotRetainTheView(t *testing.T) {
	b := []byte("mutable")
	if !EqualString(b, "mutable") {
		t.Fatal("EqualString = false, want true")
	}
	b[0] = 'M'
	if EqualString(b, "mutable") {
		t.Error("EqualString = true after the bytes changed, want false")
	}
}
