package contains

import "testing"

func TestContains(t *testing.T) {
	xs := []string{"a", "b", "c"}
	cases := []struct {
		target string
		want   bool
	}{
		{"a", true}, {"c", true}, {"z", false}, {"", false},
	}
	for _, c := range cases {
		if got := Contains(xs, c.target); got != c.want {
			t.Errorf("Contains(%v,%q)=%v; want %v", xs, c.target, got, c.want)
		}
	}
	if Contains(nil, "x") {
		t.Error("nil slice contains nothing")
	}
}
