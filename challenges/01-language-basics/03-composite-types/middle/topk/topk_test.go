package topk

import (
	"reflect"
	"testing"
)

func TestTopK(t *testing.T) {
	xs := []string{"a", "b", "a", "c", "b", "a", "d"}
	got := TopK(xs, 2)
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopK=%v; want %v", got, want)
	}
	// tie broken alphabetically: c and d both count 1
	got = TopK(xs, 4)
	want = []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopK ties=%v; want %v", got, want)
	}
}
