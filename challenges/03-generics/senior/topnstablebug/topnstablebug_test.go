package topnstablebug

import "testing"

type row struct {
	name  string
	score int
}

func score(r row) int { return r.score }

func TestTopNPicksBest(t *testing.T) {
	rows := []row{{"a", 1}, {"b", 9}, {"c", 5}}
	got := TopN(rows, score, 2)
	if len(got) != 2 || got[0].name != "b" || got[1].name != "c" {
		t.Errorf("TopN = %v, want [b c]", got)
	}
}

func TestTopNDoesNotTouchTheInput(t *testing.T) {
	rows := []row{{"a", 1}, {"b", 9}, {"c", 5}}
	TopN(rows, score, 2)
	for i, want := range []string{"a", "b", "c"} {
		if rows[i].name != want {
			t.Fatalf("input reordered: rows[%d] = %q, want %q", i, rows[i].name, want)
		}
	}
}

func TestTopNClamps(t *testing.T) {
	rows := []row{{"a", 1}}
	if got := TopN(rows, score, 99); len(got) != 1 {
		t.Errorf("TopN = %v, want 1 row", got)
	}
	if got := TopN(rows, score, -1); len(got) != 0 {
		t.Errorf("TopN = %v, want []", got)
	}
}
