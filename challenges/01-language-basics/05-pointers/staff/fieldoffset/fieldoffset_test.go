package fieldoffset

import "testing"

func TestSecondField(t *testing.T) {
	p := &Pair{A: 1, B: 2}
	if got := SecondField(p); got != 2 {
		t.Errorf("=%d want 2 (used Sizeof instead of Offsetof?)", got)
	}
}
