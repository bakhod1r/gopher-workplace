package batcherbug

import (
	"reflect"
	"testing"
)

func TestBatchesPartialTail(t *testing.T) {
	got := Batches([]int{1, 2, 3}, 2)
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {3}}) {
		t.Errorf("Batches = %v, want [[1 2] [3]]", got)
	}
}

func TestBatchesExact(t *testing.T) {
	got := Batches([]int{1, 2, 3, 4}, 2)
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {3, 4}}) {
		t.Errorf("Batches = %v, want [[1 2] [3 4]]", got)
	}
}

func TestBatchesBadSize(t *testing.T) {
	if got := Batches([]int{1}, 0); len(got) != 0 {
		t.Errorf("Batches = %v, want []", got)
	}
}
