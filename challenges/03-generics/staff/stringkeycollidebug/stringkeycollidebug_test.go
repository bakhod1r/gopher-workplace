package stringkeycollidebug

import (
	"reflect"
	"testing"
	"time"
)

func TestDistinctDropsRealDuplicates(t *testing.T) {
	if got := Distinct([]any{1, 1, 2}); !reflect.DeepEqual(got, []any{1, 2}) {
		t.Errorf("Distinct = %v, want [1 2]", got)
	}
	if got := Distinct([]string{"a", "b", "a"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Distinct = %v, want [a b]", got)
	}
}

func TestDistinctKeepsDifferentTypesThatPrintAlike(t *testing.T) {
	got := Distinct([]any{1, "1"})
	if len(got) != 2 {
		t.Fatalf("Distinct = %v, want both the int 1 and the string \"1\"", got)
	}
	if _, ok := got[0].(int); !ok {
		t.Errorf("got[0] = %#v, want the int 1", got[0])
	}
	if _, ok := got[1].(string); !ok {
		t.Errorf("got[1] = %#v, want the string \"1\"", got[1])
	}
}

func TestDistinctKeepsStructsThatPrintAlike(t *testing.T) {
	type point struct{ X, Y int }
	type pair struct{ X, Y int }
	got := Distinct([]any{point{1, 2}, pair{1, 2}})
	if len(got) != 2 {
		t.Errorf("Distinct = %v, want 2 elements: the two struct types are not equal under ==", got)
	}
}

func TestDistinctAtScale(t *testing.T) {
	const n = 200000
	vals := make([]int, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, i)
	}
	start := time.Now()
	got := Distinct(vals)
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Errorf("Distinct over %d elements took %v, ceiling 2s: every element is being formatted", n, elapsed)
	}
	if len(got) != n {
		t.Fatalf("len(Distinct) = %d, want %d", len(got), n)
	}
}
