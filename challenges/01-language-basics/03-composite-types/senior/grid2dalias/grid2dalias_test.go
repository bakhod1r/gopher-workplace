package grid2dalias

import "testing"

func TestNew(t *testing.T) {
	g := New(3, 2)
	g[0][0] = 9
	if g[1][0] != 0 || g[2][0] != 0 {
		t.Errorf("rows alias: writing g[0][0] changed others: %v", g)
	}
	if len(g) != 3 || len(g[0]) != 2 {
		t.Errorf("shape wrong: %v", g)
	}
}
