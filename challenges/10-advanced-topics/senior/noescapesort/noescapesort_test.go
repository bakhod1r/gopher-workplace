package noescapesort

import (
	"sort"
	"testing"
)

var sink int

func TestMedian3(t *testing.T) {
	cases := [][4]int{
		{1, 2, 3, 2}, {3, 2, 1, 2}, {2, 1, 3, 2},
		{1, 1, 1, 1}, {5, 5, 1, 5}, {-1, 0, 1, 0},
	}
	for _, c := range cases {
		if got := Median3(c[0], c[1], c[2]); got != c[3] {
			t.Errorf("Median3(%d,%d,%d) = %d, want %d", c[0], c[1], c[2], got, c[3])
		}
	}
}

func TestMedian3MatchesSorting(t *testing.T) {
	for a := -3; a <= 3; a++ {
		for b := -3; b <= 3; b++ {
			for c := -3; c <= 3; c++ {
				s := []int{a, b, c}
				sort.Ints(s)
				if got := Median3(a, b, c); got != s[1] {
					t.Fatalf("Median3(%d,%d,%d) = %d, want %d", a, b, c, got, s[1])
				}
			}
		}
	}
}

func TestMedian3AllocatesNothing(t *testing.T) {
	if n := testing.AllocsPerRun(500, func() { sink = Median3(9, 2, 5) }); n != 0 {
		t.Errorf("Median3 made %v allocations, want 0: no slice, no interface", n)
	}
}
