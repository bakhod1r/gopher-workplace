package buildlog

import (
	"reflect"
	"testing"
)

func TestTailReverse(t *testing.T) {
	cases := []struct {
		name  string
		lines []string
		want  []string
	}{
		{"three_lines", []string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{"single", []string{"x"}, []string{"x"}},
		{"empty_log", nil, []string{}},
		{"pair", []string{"go", "build"}, []string{"build", "go"}},
		{"blank_line", []string{"", "z"}, []string{"z", ""}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := TailReverse(tc.lines)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TailReverse(%v) = %#v, want %#v", tc.lines, got, tc.want)
			}
		})
	}
}
