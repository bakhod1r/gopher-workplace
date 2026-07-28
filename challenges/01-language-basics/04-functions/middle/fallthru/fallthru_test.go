package fallthru

import "testing"

func TestRank(t *testing.T) {
	cases := []struct {
		s    int
		want string
	}{
		{1, ""},
		{3, "bronze"},
		{6, "silver/bronze"},
		{9, "gold/silver/bronze"},
	}
	for _, c := range cases {
		if got := Rank(c.s); got != c.want {
			t.Errorf("Rank(%d)=%q want %q", c.s, got, c.want)
		}
	}
}
