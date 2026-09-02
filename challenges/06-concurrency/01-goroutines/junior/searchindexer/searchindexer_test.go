package searchindexer

import (
	"reflect"
	"testing"
)

func TestTermCounts(t *testing.T) {
	cases := []struct {
		name string
		docs []string
		term string
		want []int
	}{
		{"single_letter", []string{"banana", "abc"}, "a", []int{3, 1}},
		{"whole_word", []string{"go go"}, "go", []int{2}},
		{"absent", []string{"rust"}, "go", []int{0}},
		{"no_overlap", []string{"aaa"}, "aa", []int{1}},
		{"empty", []string{}, "a", []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TermCounts(tc.docs, tc.term); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TermCounts(%v) = %v, want %v", tc.docs, got, tc.want)
			}
		})
	}
}
