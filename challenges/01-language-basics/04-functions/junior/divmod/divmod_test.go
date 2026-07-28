package divmod

import "testing"

func TestDivMod(t *testing.T) {
	cases := []struct{ a, b, q, r int }{
		{7, 3, 2, 1},
		{10, 5, 2, 0},
		{9, 4, 2, 1},
	}
	for _, c := range cases {
		q, r := DivMod(c.a, c.b)
		if q != c.q || r != c.r {
			t.Errorf("DivMod(%d,%d)=%d,%d want %d,%d", c.a, c.b, q, r, c.q, c.r)
		}
	}
}
