package twodshallow

import "testing"

func TestClone(t *testing.T) {
	g := [][]int{{1, 2}, {3, 4}}
	c := Clone(g)
	c[0][0] = 99
	if g[0][0] != 1 {
		t.Errorf("rows shared: g=%v", g)
	}
}
