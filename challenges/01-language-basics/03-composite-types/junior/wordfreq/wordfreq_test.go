package wordfreq

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want map[string]int
	}{
		{"repeats", []string{"a", "b", "a"}, map[string]int{"a": 2, "b": 1}},
		{"single", []string{"x"}, map[string]int{"x": 1}},
		{"all same", []string{"go", "go", "go"}, map[string]int{"go": 3}},
		{"nil", nil, map[string]int{}},
		{"empty", []string{}, map[string]int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Count(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
