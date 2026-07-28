package joinmanual

import "testing"

func TestJoin(t *testing.T) {
	cases := []struct {
		parts []string
		sep   string
		want  string
	}{
		{[]string{"a", "b", "c"}, ",", "a,b,c"},
		{[]string{"x"}, ",", "x"},
		{[]string{}, ",", ""},
		{[]string{"a", "b"}, " - ", "a - b"},
	}
	for _, c := range cases {
		if got := Join(c.parts, c.sep); got != c.want {
			t.Errorf("Join(%v,%q)=%q; want %q", c.parts, c.sep, got, c.want)
		}
	}
}
