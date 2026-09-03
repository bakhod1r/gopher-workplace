package cutbytes

import (
	"testing"
	"unsafe"
)

var (
	sinkA, sinkB string
	sinkOK       bool
)

func TestCut(t *testing.T) {
	cases := []struct {
		in            string
		before, after string
		found         bool
	}{
		{"a=b", "a", "b", true},
		{"a=b=c", "a", "b=c", true},
		{"=x", "", "x", true},
		{"x=", "x", "", true},
		{"abc", "abc", "", false},
		{"", "", "", false},
		{"=", "", "", true},
	}
	for _, c := range cases {
		b, a, ok := Cut(c.in, '=')
		if b != c.before || a != c.after || ok != c.found {
			t.Errorf("Cut(%q) = %q, %q, %v, want %q, %q, %v",
				c.in, b, a, ok, c.before, c.after, c.found)
		}
	}
}

func TestCutResultsAreSubstrings(t *testing.T) {
	s := "key=value"
	before, after, ok := Cut(s, '=')
	if !ok {
		t.Fatal("Cut reported not found")
	}
	if unsafe.StringData(before) != unsafe.StringData(s) {
		t.Error("before is a copy; it must be a substring of s")
	}
	want := (*byte)(unsafe.Add(unsafe.Pointer(unsafe.StringData(s)), 4))
	if unsafe.StringData(after) != want {
		t.Error("after is a copy; it must be a substring of s")
	}
}

func TestCutAllocatesNothing(t *testing.T) {
	s := "a-very-long-configuration-key=and-an-equally-long-value"
	n := testing.AllocsPerRun(200, func() { sinkA, sinkB, sinkOK = Cut(s, '=') })
	if n != 0 {
		t.Errorf("Cut made %v allocations, want 0", n)
	}
}
