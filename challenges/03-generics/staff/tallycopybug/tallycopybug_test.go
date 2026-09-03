package tallycopybug

import "testing"

func TestBumpAllUpdatesTotals(t *testing.T) {
	ts := []Tally[string]{NewTally[string]("a"), NewTally[string]("b")}
	BumpAll(ts, "x")
	BumpAll(ts, "x")
	for i := range ts {
		if got := ts[i].Counts["x"]; got != 2 {
			t.Errorf("ts[%d].Counts[x] = %d, want 2", i, got)
		}
		if got := ts[i].Total; got != 2 {
			t.Errorf("ts[%d].Total = %d, want 2", i, got)
		}
	}
}

func TestBumpAllKeepsTalliesConsistent(t *testing.T) {
	ts := []Tally[int]{NewTally[int]("only")}
	BumpAll(ts, 1)
	BumpAll(ts, 2)
	BumpAll(ts, 1)
	if !Consistent(ts[0]) {
		t.Errorf("Consistent = false: Total=%d, Counts=%v", ts[0].Total, ts[0].Counts)
	}
	if ts[0].Total != 3 {
		t.Errorf("Total = %d, want 3", ts[0].Total)
	}
}

func TestBumpAllEmpty(t *testing.T) {
	BumpAll([]Tally[string]{}, "x")
}
