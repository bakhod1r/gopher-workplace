package factbug

import "testing"

func TestFact(t *testing.T) {
	cases := []struct{ n, want int }{{0, 1}, {1, 1}, {5, 120}, {6, 720}}
	for _, c := range cases {
		if got := Fact(c.n); got != c.want {
			t.Errorf("Fact(%d)=%d want %d", c.n, got, c.want)
		}
	}
}
