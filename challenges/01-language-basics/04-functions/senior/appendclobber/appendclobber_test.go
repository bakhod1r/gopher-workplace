package appendclobber

import "testing"

func TestTwoTails(t *testing.T) {
	base := make([]int, 2, 8) // spare capacity on purpose
	base[0], base[1] = 1, 2
	a, b := TwoTails(base, 100, 200)
	if a[len(a)-1] != 100 {
		t.Errorf("a tail=%d want 100 (clobbered by b)", a[len(a)-1])
	}
	if b[len(b)-1] != 200 {
		t.Errorf("b tail=%d want 200", b[len(b)-1])
	}
}
