package splitmanual

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		s    string
		sep  byte
		want []string
	}{
		{"a,b,c", ',', []string{"a", "b", "c"}},
		{"a,,c", ',', []string{"a", "", "c"}},
		{"abc", ',', []string{"abc"}},
		{"", ',', []string{""}},
		{",", ',', []string{"", ""}},
	}
	for _, c := range cases {
		got := Split(c.s, c.sep)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Split(%q,%q)=%v; want %v", c.s, c.sep, got, c.want)
		}
	}
}
