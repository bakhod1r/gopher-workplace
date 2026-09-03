package bytesandrunes

import "testing"

var sinkA, sinkB int

func TestCounts(t *testing.T) {
	cases := []struct {
		in           string
		bytes, runes int
	}{
		{"hello", 5, 5},
		{"héllo", 6, 5},
		{"日本語", 9, 3},
		{"", 0, 0},
		{"a\u00e9\u65e5", 6, 3},
	}
	for _, c := range cases {
		b, r := Counts(c.in)
		if b != c.bytes || r != c.runes {
			t.Errorf("Counts(%q) = %d, %d, want %d, %d", c.in, b, r, c.bytes, c.runes)
		}
	}
}

func TestCountsASCIIAgree(t *testing.T) {
	b, r := Counts("plain ascii")
	if b != r {
		t.Errorf("bytes = %d, runes = %d, want them equal for ASCII", b, r)
	}
}

func TestCountsAllocatesNothing(t *testing.T) {
	s := "a string with sömé nön-ascii characters in it"
	if n := testing.AllocsPerRun(200, func() { sinkA, sinkB = Counts(s) }); n != 0 {
		t.Errorf("Counts made %v allocations, want 0: do not convert to []rune", n)
	}
}
