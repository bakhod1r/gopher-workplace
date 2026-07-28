package matmul

import (
	"reflect"
	"testing"
)

func TestMul(t *testing.T) {
	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{5, 6}, {7, 8}}
	got := Mul(a, b)
	want := [][]int{{19, 22}, {43, 50}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mul=%v; want %v", got, want)
	}
	if Mul([][]int{{1, 2, 3}}, [][]int{{1}}) != nil {
		t.Error("incompatible -> nil")
	}
}
