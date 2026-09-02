package batchergen

import (
	"reflect"
	"testing"
)

func TestBatcherFills(t *testing.T) {
	b := NewBatcher[int](2)
	if got, ok := b.Add(1); ok || got != nil {
		t.Errorf("Add(1) = %v, %v, want nil, false", got, ok)
	}
	got, ok := b.Add(2)
	if !ok || !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Add(2) = %v, %v, want [1 2], true", got, ok)
	}
	if got, ok := b.Add(3); ok {
		t.Errorf("Add(3) = %v, %v, want nil, false (buffer must reset)", got, ok)
	}
}

func TestBatcherFlush(t *testing.T) {
	b := NewBatcher[int](3)
	b.Add(1)
	got, ok := b.Flush()
	if !ok || !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Flush = %v, %v, want [1], true", got, ok)
	}
	if _, ok := b.Flush(); ok {
		t.Error("Flush on an empty batcher = true, want false")
	}
}

func TestBatcherMinimumSize(t *testing.T) {
	b := NewBatcher[int](0)
	got, ok := b.Add(1)
	if !ok || !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Add = %v, %v, want [1], true (size below 1 means 1)", got, ok)
	}
}
