package lowerboundbug

import (
	"testing"
	"time"
)

func TestLowerBoundDuplicates(t *testing.T) {
	s := []int{1, 2, 2, 2, 3}
	if got := LowerBound(s, 2); got != 1 {
		t.Errorf("LowerBound = %d, want 1 (first element >= 2)", got)
	}
	if got := LowerBound(s, 3); got != 4 {
		t.Errorf("LowerBound = %d, want 4", got)
	}
}

func TestLowerBoundEdges(t *testing.T) {
	if got := LowerBound([]int{}, 5); got != 0 {
		t.Errorf("LowerBound = %d, want 0", got)
	}
	if got := LowerBound([]int{1, 3}, 0); got != 0 {
		t.Errorf("LowerBound = %d, want 0", got)
	}
	if got := LowerBound([]int{1, 3}, 2); got != 1 {
		t.Errorf("LowerBound = %d, want 1", got)
	}
	if got := LowerBound([]int{1, 3}, 9); got != 2 {
		t.Errorf("LowerBound = %d, want 2", got)
	}
}

func TestLowerBoundStrings(t *testing.T) {
	s := []string{"a", "b", "b", "d"}
	if got := LowerBound(s, "b"); got != 1 {
		t.Errorf("LowerBound = %d, want 1", got)
	}
	if got := LowerBound(s, "c"); got != 3 {
		t.Errorf("LowerBound = %d, want 3", got)
	}
}

func TestLowerBoundIsLogarithmic(t *testing.T) {
	const n = 4000000
	const queries = 200000
	const budget = 2 * time.Second

	s := make([]int, n)
	for i := range s {
		s[i] = i * 2
	}

	start := time.Now()
	acc := 0
	for q := 0; q < queries; q++ {
		acc += LowerBound(s, (q*7919)%(2*n))
	}
	elapsed := time.Since(start)

	if acc == 0 {
		t.Fatal("no query produced an index")
	}
	if elapsed > budget {
		t.Errorf("%d queries over %d elements took %v, budget %v: the search is not logarithmic",
			queries, n, elapsed, budget)
	}
}
