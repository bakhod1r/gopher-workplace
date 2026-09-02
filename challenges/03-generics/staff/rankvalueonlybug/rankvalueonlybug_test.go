package rankvalueonlybug

import (
	"reflect"
	"testing"
	"time"
)

func TestRankByValue(t *testing.T) {
	got := Rank(map[string]int{"a": 2, "b": 1})
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Rank = %v, want [a b]", got)
	}
}

func TestRankBreaksTiesByKey(t *testing.T) {
	m := map[string]int{"h": 1, "g": 1, "f": 1, "e": 1, "d": 1, "c": 1, "b": 1, "a": 1}
	want := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for i := 0; i < 50; i++ {
		if got := Rank(m); !reflect.DeepEqual(got, want) {
			t.Fatalf("run %d: Rank = %v, want %v", i, got, want)
		}
	}
}

func TestRankEmpty(t *testing.T) {
	if got := Rank(map[string]int{}); len(got) != 0 {
		t.Errorf("Rank = %v, want []", got)
	}
}

func TestRankScale(t *testing.T) {
	const n = 200_000
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[i] = i % 4
	}
	start := time.Now()
	got := Rank(m)
	elapsed := time.Since(start)
	if len(got) != n {
		t.Fatalf("len(Rank) = %d, want %d", len(got), n)
	}
	for i := 1; i < len(got); i++ {
		a, b := got[i-1], got[i]
		if m[a] < m[b] || (m[a] == m[b] && a > b) {
			t.Fatalf("not a total order at %d: %d(%d) before %d(%d)", i, a, m[a], b, m[b])
		}
	}
	if elapsed > 2*time.Second {
		t.Errorf("Rank of %d keys took %v, want under 2s", n, elapsed)
	}
}
