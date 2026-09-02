package slugifier

import (
	"reflect"
	"testing"
)

func TestSlugs(t *testing.T) {
	cases := []struct {
		name   string
		titles []string
		want   []string
	}{
		{"two_words", []string{"Hello World"}, []string{"hello-world"}},
		{"blank", []string{"   "}, []string{""}},
		{"empty", []string{}, []string{}},
		{"collapses_runs", []string{"Go   Is  Fast"}, []string{"go-is-fast"}},
		{"batch", []string{"A B", "C"}, []string{"a-b", "c"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Slugs(tc.titles); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Slugs(%v) = %v, want %v", tc.titles, got, tc.want)
			}
		})
	}
}
