package transposegen

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	if got := Transpose([][]int{{1, 2}, {3, 4}}); !reflect.DeepEqual(got, [][]int{{1, 3}, {2, 4}}) {
		t.Errorf("Transpose = %v, want [[1 3] [2 4]]", got)
	}
	if got := Transpose([][]int{{1, 2, 3}}); !reflect.DeepEqual(got, [][]int{{1}, {2}, {3}}) {
		t.Errorf("Transpose = %v, want [[1] [2] [3]]", got)
	}
}

func TestTransposeRejectsRagged(t *testing.T) {
	if got := Transpose([][]int{{1, 2}, {3}}); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Transpose(ragged) = %v, want []", got)
	}
}

func TestTransposeEdges(t *testing.T) {
	if got := Transpose([][]int{}); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Transpose(empty) = %v, want []", got)
	}
	if got := Transpose([][]int{{}, {}}); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Transpose(zero width) = %v, want []", got)
	}
	if got := Transpose([][]string{{"a"}, {"b"}}); !reflect.DeepEqual(got, [][]string{{"a", "b"}}) {
		t.Errorf("Transpose = %v, want [[a b]]", got)
	}
}
