package clone2dgen

import (
	"reflect"
	"testing"
)

func TestClone2D(t *testing.T) {
	m := [][]int{{1, 2}, {3}}
	got := Clone2D(m)
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {3}}) {
		t.Errorf("Clone2D = %v, want [[1 2] [3]]", got)
	}
}

func TestClone2DRowsAreIndependent(t *testing.T) {
	m := [][]int{{1, 2}}
	got := Clone2D(m)
	got[0][0] = 99
	if m[0][0] != 1 {
		t.Errorf("writing into a copied row changed the original: %v", m)
	}
}

func TestClone2DEdges(t *testing.T) {
	got := Clone2D([][]string(nil))
	if got == nil || len(got) != 0 {
		t.Errorf("Clone2D(nil) = %v, want an empty non-nil result", got)
	}
	rows := Clone2D([][]int{{}})
	if len(rows) != 1 || len(rows[0]) != 0 {
		t.Errorf("Clone2D([[]]) = %v, want one empty row", rows)
	}
}
