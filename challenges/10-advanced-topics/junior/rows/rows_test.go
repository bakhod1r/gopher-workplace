package rows

import "testing"

func TestRowsShape(t *testing.T) {
	g := Rows(2, 3)
	if len(g) != 2 {
		t.Fatalf("rows = %d, want 2", len(g))
	}
	for i, row := range g {
		if len(row) != 3 {
			t.Fatalf("row %d has length %d, want 3", i, len(row))
		}
		for _, v := range row {
			if v != 0 {
				t.Fatalf("row %d is not zeroed", i)
			}
		}
	}
	if g := Rows(0, 3); g != nil {
		t.Errorf("Rows(0,3) = %v, want nil", g)
	}
}

func TestRowsShareOneArray(t *testing.T) {
	g := Rows(2, 3)
	if &g[0][0] == &g[1][0] {
		t.Fatal("rows overlap")
	}
	g[0] = append(g[0], 7)
	if g[1][0] == 7 {
		t.Error("row 0 spilled into row 1: cap each row with a three-index slice")
	}
}

func TestRowsAllocatesTwice(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { _ = Rows(32, 32) }); n > 2 {
		t.Errorf("Rows made %v allocations, want at most 2", n)
	}
}
