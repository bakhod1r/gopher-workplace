package chunk

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunk = %v, want %v", got, want)
	}
	if got := Chunk([]int{1, 2}, 5); !reflect.DeepEqual(got, [][]int{{1, 2}}) {
		t.Errorf("Chunk = %v, want [[1 2]]", got)
	}
	if got := Chunk(nil, 3); len(got) != 0 {
		t.Errorf("Chunk(nil) = %v, want empty", got)
	}
	if got := Chunk([]int{1}, 0); got != nil {
		t.Errorf("Chunk(s, 0) = %v, want nil", got)
	}
}

func TestChunkGroupsAreViews(t *testing.T) {
	s := []int{1, 2, 3, 4}
	g := Chunk(s, 2)
	g[0][0] = 99
	if s[0] != 99 {
		t.Error("the groups copied the elements; they must be views into s")
	}
}
