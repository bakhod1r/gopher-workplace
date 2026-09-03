package deferloop

import (
	"reflect"
	"testing"
)

func TestProcessDoubles(t *testing.T) {
	got := Process([]int{1, 2, 3}, func(int) {})
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("Process = %v, want [2 4 6]", got)
	}
}

func TestReleaseHappensPerItem(t *testing.T) {
	held := 0
	peak := 0
	Process([]int{1, 2, 3, 4, 5}, func(int) { held-- })
	_ = peak

	held = 0
	Process([]int{1, 2, 3, 4, 5}, func(int) {
		held--
	})
	if held != -5 {
		t.Fatalf("release called %d times, want 5", -held)
	}
}

func TestReleaseOrderIsForward(t *testing.T) {
	var order []int
	Process([]int{1, 2, 3}, func(v int) { order = append(order, v) })
	if !reflect.DeepEqual(order, []int{1, 2, 3}) {
		t.Errorf("release order = %v, want [1 2 3]: each item must be released as it finishes", order)
	}
}

func TestReleaseIsNotStackedToTheEnd(t *testing.T) {
	outstanding := 0
	max := 0
	Process([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(int) { outstanding-- })
	_ = max
	if outstanding != -8 {
		t.Fatalf("release called %d times, want 8", -outstanding)
	}

	seen := make([]int, 0, 8)
	Process([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(v int) { seen = append(seen, v) })
	for i, v := range seen {
		if v != i+1 {
			t.Fatalf("release sequence = %v, want ascending: the cleanups piled up until the return", seen)
		}
	}
}
