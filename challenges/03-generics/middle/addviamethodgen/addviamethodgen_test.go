package addviamethodgen

import "testing"

func TestSumAll(t *testing.T) {
	got := SumAll([]Money{{200}, {350}})
	if got.Cents != 550 {
		t.Errorf("SumAll = %+v, want 550 cents", got)
	}
}

func TestSumAllEdges(t *testing.T) {
	if got := SumAll([]Money{}); got.Cents != 0 {
		t.Errorf("SumAll(empty) = %+v, want 0", got)
	}
	if got := SumAll([]Money{{7}}); got.Cents != 7 {
		t.Errorf("SumAll(single) = %+v, want 7", got)
	}
	if got := SumAll([]Money{{-5}, {5}}); got.Cents != 0 {
		t.Errorf("SumAll = %+v, want 0", got)
	}
}
