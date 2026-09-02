package chunkgen

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	if got := Chunk([]int{1, 2, 3}, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {3}}) {
		t.Errorf("Chunk([]int{1, 2, 3}, 2) = %v, want [[1 2] [3]]", got)
	}
	if got := Chunk([]int{1, 2, 3, 4}, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {3, 4}}) {
		t.Errorf("Chunk([]int{1, 2, 3, 4}, 2) = %v, want [[1 2] [3 4]]", got)
	}
	if got := Chunk([]int{1, 2}, 5); !reflect.DeepEqual(got, [][]int{{1, 2}}) {
		t.Errorf("Chunk([]int{1, 2}, 5) = %v, want [[1 2]]", got)
	}
	if got := Chunk([]int{}, 2); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Chunk([]int{}, 2) = %v, want []", got)
	}
	if got := Chunk([]int{1}, 0); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Chunk([]int{1}, 0) = %v, want []", got)
	}
}

func TestChunkGroupsAreIndependent(t *testing.T) {
	in := []int{1, 2, 3, 4}
	got := Chunk(in, 2)
	got[0] = append(got[0], 99)
	if !reflect.DeepEqual(in, []int{1, 2, 3, 4}) {
		t.Errorf("appending to a group mutated the input: %v, want [1 2 3 4]", in)
	}
}
