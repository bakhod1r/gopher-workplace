package skipunexported

import "testing"

type mix struct {
	A      int
	B      int
	hidden int
	Name   string
}

func TestSumInts(t *testing.T) {
	if got := SumInts(mix{A: 1, B: 2, hidden: 100, Name: "x"}); got != 3 {
		t.Errorf("SumInts = %d, want 3: only exported int fields count", got)
	}
	if got := SumInts(mix{}); got != 0 {
		t.Errorf("SumInts = %d, want 0", got)
	}
}

func TestSumIntsNegative(t *testing.T) {
	if got := SumInts(mix{A: -5, B: 5}); got != 0 {
		t.Errorf("SumInts = %d, want 0", got)
	}
}

func TestSumIntsRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &mix{}} {
		if got := SumInts(in); got != 0 {
			t.Errorf("SumInts(%#v) = %d, want 0", in, got)
		}
	}
}

func TestSumIntsAllExported(t *testing.T) {
	type open struct{ A, B, C int }
	if got := SumInts(open{1, 2, 3}); got != 6 {
		t.Errorf("SumInts = %d, want 6", got)
	}
}
