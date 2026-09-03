package clone2dsharebug

import (
	"reflect"
	"testing"
)

func TestClone2DCopiesValues(t *testing.T) {
	m := [][]int{{1, 2}, {3}}
	got := Clone2D(m)
	if !reflect.DeepEqual(got, m) {
		t.Errorf("Clone2D = %v, want %v", got, m)
	}
}

func TestClone2DRowsAreIndependent(t *testing.T) {
	m := [][]int{{1, 2}, {3, 4}}
	c := Clone2D(m)
	c[0][0] = 99
	c[1][1] = 98
	if m[0][0] != 1 || m[1][1] != 4 {
		t.Errorf("source mutated: %v, want [[1 2] [3 4]]", m)
	}
}

func TestClone2DEmpty(t *testing.T) {
	if got := Clone2D[int](nil); len(got) != 0 {
		t.Errorf("Clone2D = %v, want []", got)
	}
}
