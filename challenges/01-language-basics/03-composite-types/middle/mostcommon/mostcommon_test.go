package mostcommon

import "testing"

func TestMode(t *testing.T) {
	cases := []struct {
		xs   []int
		want int
		ok   bool
	}{
		{[]int{1, 2, 2, 3}, 2, true},
		{[]int{1, 1, 2, 2}, 1, true}, // tie -> smaller
		{[]int{5}, 5, true},
		{nil, 0, false},
	}
	for _, c := range cases {
		got, ok := Mode(c.xs)
		if got != c.want || ok != c.ok {
			t.Errorf("Mode(%v)=(%d,%v); want (%d,%v)", c.xs, got, ok, c.want, c.ok)
		}
	}
}
