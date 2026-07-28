package classifybug

import "testing"

func TestClass(t *testing.T) {
	cases := []struct {
		c    int
		want string
	}{{200, "success"}, {404, "client"}, {500, "server"}, {301, "unknown"}, {100, "unknown"}}
	for _, c := range cases {
		if got := Class(c.c); got != c.want {
			t.Errorf("Class(%d)=%q want %q", c.c, got, c.want)
		}
	}
}
