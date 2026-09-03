package prependquadbug

import (
	"reflect"
	"testing"
	"time"
)

func TestReversedOrder(t *testing.T) {
	if got := Reversed([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("Reversed = %v, want [3 2 1]", got)
	}
	if got := Reversed([]string{"a", "b"}); !reflect.DeepEqual(got, []string{"b", "a"}) {
		t.Errorf("Reversed = %v, want [b a]", got)
	}
}

func TestReversedEmptyAndSingle(t *testing.T) {
	if got := Reversed([]int{}); len(got) != 0 {
		t.Errorf("Reversed = %v, want []", got)
	}
	if got := Reversed([]int{7}); !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("Reversed = %v, want [7]", got)
	}
}

func TestReversedDoesNotTouchInput(t *testing.T) {
	in := []int{1, 2, 3}
	_ = Reversed(in)
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input = %v, want [1 2 3]", in)
	}
}

func TestReversedIsLinear(t *testing.T) {
	const n = 120000
	const budget = 2 * time.Second

	in := make([]int, n)
	for i := range in {
		in[i] = i
	}

	start := time.Now()
	got := Reversed(in)
	elapsed := time.Since(start)

	if len(got) != n || got[0] != n-1 || got[n-1] != 0 {
		t.Fatalf("Reversed produced the wrong result for n=%d", n)
	}
	if elapsed > budget {
		t.Errorf("Reversed(%d elements) took %v, budget %v: the implementation is not linear", n, elapsed, budget)
	}
}
