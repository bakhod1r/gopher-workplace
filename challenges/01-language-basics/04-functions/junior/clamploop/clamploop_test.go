package clamploop

import (
	"reflect"
	"testing"
)

func TestClampAll(t *testing.T) {
	in := []int{-5, 3, 20}
	got := ClampAll(in, 0, 10)
	if !reflect.DeepEqual(got, []int{0, 3, 10}) {
		t.Errorf("=%v want [0 3 10]", got)
	}
	if !reflect.DeepEqual(in, []int{-5, 3, 20}) {
		t.Errorf("input mutated: %v", in)
	}
}
