package logshipper

import (
	"reflect"
	"testing"
)

func TestPayloadSizes(t *testing.T) {
	cases := []struct {
		name  string
		lines []string
		want  []int
	}{
		{"two_lines", []string{"ok", "boom"}, []int{3, 5}},
		{"blank_line", []string{""}, []int{1}},
		{"empty", []string{}, []int{}},
		{"unicode_counts_bytes", []string{"añ"}, []int{4}},
		{"batch", []string{"a", "bb", "ccc"}, []int{2, 3, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PayloadSizes(tc.lines); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PayloadSizes(%v) = %v, want %v", tc.lines, got, tc.want)
			}
		})
	}
}
