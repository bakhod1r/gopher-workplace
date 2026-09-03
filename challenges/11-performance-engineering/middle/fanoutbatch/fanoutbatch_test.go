package fanoutbatch

import (
	"reflect"
	"testing"
)

func TestChunks(t *testing.T) {
	got := Chunks([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunks = %v, want %v", got, want)
	}
}

func TestChunksExactMultiple(t *testing.T) {
	got := Chunks([]int{1, 2, 3, 4}, 2)
	want := [][]int{{1, 2}, {3, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunks = %v, want %v", got, want)
	}
}

func TestChunksGuards(t *testing.T) {
	for _, c := range []struct {
		items []int
		size  int
	}{{[]int{1, 2}, 0}, {[]int{1, 2}, -1}, {nil, 2}} {
		got := Chunks(c.items, c.size)
		if got == nil || len(got) != 0 {
			t.Errorf("Chunks(%v, %d) = %v, want empty non-nil slice", c.items, c.size, got)
		}
	}
}

func TestChunksShareTheInput(t *testing.T) {
	items := []int{1, 2, 3, 4}
	got := Chunks(items, 2)
	got[0][0] = 99
	if items[0] != 99 {
		t.Error("chunks copied the data; they should be sub-slices of the input")
	}
}

func TestSumBatches(t *testing.T) {
	if got := SumBatches([]int{1, 2, 3, 4, 5}, 2); !reflect.DeepEqual(got, []int{3, 7, 5}) {
		t.Errorf("SumBatches = %v, want [3 7 5]", got)
	}
	if got := SumBatches(nil, 2); got == nil || len(got) != 0 {
		t.Errorf("SumBatches(nil) = %v, want empty non-nil slice", got)
	}
	if got := SumBatches([]int{1, 2}, 0); got == nil || len(got) != 0 {
		t.Errorf("SumBatches(_, 0) = %v, want empty non-nil slice", got)
	}
}

func TestSumBatchesIsDeterministic(t *testing.T) {
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}
	first := SumBatches(items, 7)
	for i := 0; i < 20; i++ {
		if !reflect.DeepEqual(SumBatches(items, 7), first) {
			t.Fatal("SumBatches results moved between runs; batch order must be preserved")
		}
	}
	var total int
	for _, s := range first {
		total += s
	}
	if total != 999*1000/2 {
		t.Errorf("total = %d, want %d", total, 999*1000/2)
	}
}
