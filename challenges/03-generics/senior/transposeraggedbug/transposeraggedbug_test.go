package transposeraggedbug

import (
	"reflect"
	"testing"
)

func TestTransposeFirstRowShort(t *testing.T) {
	got := Transpose([][]int{{1}, {2, 3}})
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {3}}) {
		t.Errorf("Transpose = %v, want [[1 2] [3]]", got)
	}
}

func TestTransposeFirstRowLong(t *testing.T) {
	got := Transpose([][]int{{1, 2}, {3}})
	if !reflect.DeepEqual(got, [][]int{{1, 3}, {2}}) {
		t.Errorf("Transpose = %v, want [[1 3] [2]]", got)
	}
}

func TestTransposeEmpty(t *testing.T) {
	if got := Transpose([][]int{}); len(got) != 0 {
		t.Errorf("Transpose = %v, want []", got)
	}
}
