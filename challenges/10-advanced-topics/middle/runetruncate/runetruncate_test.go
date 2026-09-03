package runetruncate

import (
	"testing"
	"unicode/utf8"
)

var sink string

func TestTruncateASCII(t *testing.T) {
	if got := Truncate("hello", 3); got != "hel" {
		t.Errorf("Truncate = %q, want \"hel\"", got)
	}
	if got := Truncate("hi", 9); got != "hi" {
		t.Errorf("Truncate = %q, want \"hi\"", got)
	}
	if got := Truncate("hi", 0); got != "" {
		t.Errorf("Truncate = %q, want empty", got)
	}
	if got := Truncate("hi", -1); got != "" {
		t.Errorf("Truncate = %q, want empty", got)
	}
}

func TestTruncateMultiByte(t *testing.T) {
	// "héllo": h is 1 byte, é is 2 bytes at offsets 1-2
	s := "héllo"
	cases := map[int]string{1: "h", 2: "h", 3: "hé", 4: "hél"}
	for n, want := range cases {
		if got := Truncate(s, n); got != want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", s, n, got, want)
		}
	}
}

func TestTruncateAlwaysValid(t *testing.T) {
	s := "日本語テキスト"
	for n := 0; n <= len(s)+2; n++ {
		got := Truncate(s, n)
		if !utf8.ValidString(got) {
			t.Fatalf("Truncate(%d) = %q, which is not valid UTF-8", n, got)
		}
		if len(got) > n && n <= len(s) {
			t.Fatalf("Truncate(%d) returned %d bytes", n, len(got))
		}
	}
}

func TestTruncateAllocatesNothing(t *testing.T) {
	s := "a fairly long string with some ünicode in it"
	if n := testing.AllocsPerRun(200, func() { sink = Truncate(s, 20) }); n != 0 {
		t.Errorf("Truncate made %v allocations, want 0: return a substring", n)
	}
}
