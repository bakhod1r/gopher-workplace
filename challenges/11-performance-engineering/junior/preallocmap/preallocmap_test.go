package preallocmap

import (
	"reflect"
	"strconv"
	"testing"
)

var sink map[string]int

func TestIndexFirstOccurrence(t *testing.T) {
	got := Index([]string{"a", "b", "a", "c", "b"})
	want := map[string]int{"a": 0, "b": 1, "c": 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Index = %v, want %v", got, want)
	}
}

func TestIndexEmpty(t *testing.T) {
	got := Index(nil)
	if got == nil {
		t.Fatal("Index(nil) = nil, want empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("Index(nil) = %v, want empty", got)
	}
}

func TestIndexIsPreallocated(t *testing.T) {
	words := make([]string, 1000)
	for i := range words {
		words[i] = strconv.Itoa(i)
	}
	allocs := testing.AllocsPerRun(20, func() { sink = Index(words) })
	if allocs > 8 {
		t.Errorf("Index(1000 words) made %v allocations, want at most 8 — the size is known up front", allocs)
	}
}
