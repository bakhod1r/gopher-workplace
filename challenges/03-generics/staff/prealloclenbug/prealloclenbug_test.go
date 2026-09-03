package prealloclenbug

import (
	"reflect"
	"testing"
	"time"
)

func double(n int) int { return n * 2 }

func TestMapValues(t *testing.T) {
	got := Map([]int{1, 2}, double)
	if !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("Map = %v, want [2 4]", got)
	}
}

func TestMapEmpty(t *testing.T) {
	if got := Map([]int{}, double); len(got) != 0 {
		t.Errorf("Map = %v, want []", got)
	}
}

func TestMapChangesType(t *testing.T) {
	got := Map([]int{1, 2}, func(n int) string {
		if n == 1 {
			return "a"
		}
		return "b"
	})
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Map = %v, want [a b]", got)
	}
}

func TestMapScale(t *testing.T) {
	const n = 2_000_000
	in := make([]int, n)
	for i := range in {
		in[i] = i
	}
	start := time.Now()
	got := Map(in, double)
	elapsed := time.Since(start)
	if len(got) != n {
		t.Fatalf("len(Map) = %d, want %d", len(got), n)
	}
	if got[0] != 0 || got[n-1] != 2*(n-1) {
		t.Fatalf("Map = [%d ... %d], want [0 ... %d]", got[0], got[n-1], 2*(n-1))
	}
	if elapsed > 500*time.Millisecond {
		t.Errorf("Map over %d elements took %v, want under 500ms", n, elapsed)
	}
}
