package anagramgroup

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		words []string
		want  int
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, 3},
		{[]string{"a", "b", "c"}, 3},
		{[]string{"listen", "silent"}, 1},
		{nil, 0},
	}
	for _, c := range cases {
		if got := Count(c.words); got != c.want {
			t.Errorf("Count(%v)=%d; want %d", c.words, got, c.want)
		}
	}
}
