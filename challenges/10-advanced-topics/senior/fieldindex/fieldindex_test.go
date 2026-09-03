package fieldindex

import (
	"errors"
	"testing"
)

type rec struct {
	N      int
	M      int
	Label  string
	hidden int
}

func TestSumColumn(t *testing.T) {
	rows := []rec{{N: 1, M: 10}, {N: 2, M: 20}, {N: 3, M: 30}}
	if got, err := SumColumn(rows, "N"); err != nil || got != 6 {
		t.Errorf("SumColumn(N) = %d, %v, want 6, nil", got, err)
	}
	if got, err := SumColumn(rows, "M"); err != nil || got != 60 {
		t.Errorf("SumColumn(M) = %d, %v, want 60, nil", got, err)
	}
}

func TestSumColumnEmpty(t *testing.T) {
	if got, err := SumColumn([]rec{}, "N"); err != nil || got != 0 {
		t.Errorf("SumColumn = %d, %v, want 0, nil", got, err)
	}
}

func TestSumColumnBadShape(t *testing.T) {
	cases := []struct {
		name  string
		rows  any
		field string
	}{
		{"not a slice", rec{}, "N"},
		{"nil", nil, "N"},
		{"slice of ints", []int{1}, "N"},
		{"missing field", []rec{{}}, "Nope"},
		{"wrong kind", []rec{{}}, "Label"},
		{"unexported", []rec{{}}, "hidden"},
	}
	for _, c := range cases {
		if _, err := SumColumn(c.rows, c.field); !errors.Is(err, ErrShape) {
			t.Errorf("%s: err = %v, want ErrShape", c.name, err)
		}
	}
}

func TestSumColumnResolvesTheFieldOnce(t *testing.T) {
	rows := make([]rec, 4096)
	for i := range rows {
		rows[i].N = 1
	}
	got, err := SumColumn(rows, "N")
	if err != nil || got != 4096 {
		t.Fatalf("SumColumn = %d, %v, want 4096, nil", got, err)
	}
	n := testing.AllocsPerRun(20, func() { _, _ = SumColumn(rows, "N") })
	if n > 4 {
		t.Errorf("SumColumn made %v allocations for 4096 rows, want a handful: resolve the field once", n)
	}
}

func BenchmarkSumColumn(b *testing.B) {
	rows := make([]rec, 4096)
	for i := range rows {
		rows[i].N = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = SumColumn(rows, "N")
	}
}
