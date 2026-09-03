package stdbinsearchcmpbug

import (
	"testing"
	"time"
)

func board() []Entry {
	return []Entry{{"a", 9}, {"b", 5}, {"c", 1}}
}

func TestFindScoreMiddle(t *testing.T) {
	if i, ok := FindScore(board(), 5); !ok || i != 1 {
		t.Errorf("FindScore(5) = %d, %v, want 1, true", i, ok)
	}
}

func TestFindScoreEnds(t *testing.T) {
	if i, ok := FindScore(board(), 9); !ok || i != 0 {
		t.Errorf("FindScore(9) = %d, %v, want 0, true", i, ok)
	}
	if i, ok := FindScore(board(), 1); !ok || i != 2 {
		t.Errorf("FindScore(1) = %d, %v, want 2, true", i, ok)
	}
}

func TestFindScoreMissing(t *testing.T) {
	if i, ok := FindScore(board(), 7); ok || i != -1 {
		t.Errorf("FindScore(7) = %d, %v, want -1, false", i, ok)
	}
	if i, ok := FindScore(nil, 1); ok || i != -1 {
		t.Errorf("FindScore(nil) = %d, %v, want -1, false", i, ok)
	}
}

func TestFindScoreAtScale(t *testing.T) {
	const n = 200000
	big := make([]Entry, n)
	for i := 0; i < n; i++ {
		big[i] = Entry{Score: (n - i) * 2}
	}
	start := time.Now()
	for i := 0; i < n; i += 97 {
		got, ok := FindScore(big, (n-i)*2)
		if !ok || got != i {
			t.Fatalf("FindScore at %d = %d, %v, want %d, true", i, got, ok, i)
		}
	}
	for i := 0; i < n; i += 997 {
		if got, ok := FindScore(big, (n-i)*2-1); ok || got != -1 {
			t.Fatalf("odd score %d = %d, %v, want -1, false", i, got, ok)
		}
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("scale lookups took %v, want under 2s", d)
	}
}
