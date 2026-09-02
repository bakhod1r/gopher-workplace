package chunkbygen

import (
	"reflect"
	"testing"
)

func TestChunkBy(t *testing.T) {
	equal := func(a, b int) bool { return a == b }
	if got := ChunkBy([]int{1, 1, 2}, equal); !reflect.DeepEqual(got, [][]int{{1, 1}, {2}}) {
		t.Errorf("ChunkBy = %v, want [[1 1] [2]]", got)
	}
	always := func(a, b int) bool { return true }
	if got := ChunkBy([]int{1, 2, 3}, always); !reflect.DeepEqual(got, [][]int{{1, 2, 3}}) {
		t.Errorf("ChunkBy = %v, want [[1 2 3]]", got)
	}
	never := func(a, b int) bool { return false }
	if got := ChunkBy([]int{1, 2}, never); !reflect.DeepEqual(got, [][]int{{1}, {2}}) {
		t.Errorf("ChunkBy = %v, want [[1] [2]]", got)
	}
}

func TestChunkByEdges(t *testing.T) {
	equal := func(a, b int) bool { return a == b }
	if got := ChunkBy([]int{}, equal); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("ChunkBy(empty) = %v, want []", got)
	}
	if got := ChunkBy([]int{5}, equal); !reflect.DeepEqual(got, [][]int{{5}}) {
		t.Errorf("ChunkBy(one) = %v, want [[5]]", got)
	}
}

func TestChunkByAscendingRuns(t *testing.T) {
	rising := func(a, b int) bool { return b > a }
	got := ChunkBy([]int{1, 2, 1, 3}, rising)
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {1, 3}}) {
		t.Errorf("ChunkBy = %v, want [[1 2] [1 3]]", got)
	}
}
