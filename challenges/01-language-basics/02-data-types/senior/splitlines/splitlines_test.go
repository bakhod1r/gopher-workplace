package splitlines

import (
	"reflect"
	"testing"
)

func TestLines(t *testing.T) {
	cases := []struct {
		s    string
		want []string
	}{
		{"a\nb\nc", []string{"a", "b", "c"}},
		{"a\r\nb\r\nc", []string{"a", "b", "c"}},
		{"only", []string{"only"}},
		{"", []string{""}},
	}
	for _, c := range cases {
		got := Lines(c.s)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Lines(%q)=%q; want %q", c.s, got, c.want)
		}
	}
}
