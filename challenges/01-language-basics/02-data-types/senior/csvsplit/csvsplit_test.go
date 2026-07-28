package csvsplit

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		line string
		want []string
	}{
		{"a,b,c", []string{"a", "b", "c"}},
		{`a,"b,c",d`, []string{"a", "b,c", "d"}},
		{`"x""y",z`, []string{`x"y`, "z"}},
		{"", []string{""}},
	}
	for _, c := range cases {
		got := Split(c.line)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Split(%q)=%v; want %v", c.line, got, c.want)
		}
	}
}
