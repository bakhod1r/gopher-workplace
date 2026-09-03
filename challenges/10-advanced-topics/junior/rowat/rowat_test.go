package rowat

import "testing"

func TestRow(t *testing.T) {
	g := [][]int{{1, 2}, {3}}
	got, ok := Row(g, 0)
	if !ok || len(got) != 2 || got[0] != 1 {
		t.Errorf("Row = %v, %v, want [1 2], true", got, ok)
	}
	got, ok = Row(g, 1)
	if !ok || len(got) != 1 || got[0] != 3 {
		t.Errorf("Row = %v, %v, want [3], true", got, ok)
	}
}

func TestRowOutOfRange(t *testing.T) {
	g := [][]int{{1}}
	for _, i := range []int{-1, 1, 99} {
		if _, ok := Row(g, i); ok {
			t.Errorf("Row(_, %d) reported ok, want false", i)
		}
	}
	if _, ok := Row(nil, 0); ok {
		t.Error("Row(nil, 0) reported ok, want false")
	}
}

func TestRowIsAView(t *testing.T) {
	g := [][]int{{1, 2}}
	row, _ := Row(g, 0)
	row[0] = 99
	if g[0][0] != 99 {
		t.Error("the row is a copy; it must be a view into g")
	}
}

func TestRowNilRow(t *testing.T) {
	g := [][]int{nil}
	got, ok := Row(g, 0)
	if !ok {
		t.Error("a nil row still exists")
	}
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}
