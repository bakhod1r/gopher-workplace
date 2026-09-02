package spamfilter

import (
	"reflect"
	"testing"
)

func TestFlagged(t *testing.T) {
	cases := []struct {
		name     string
		messages []string
		banned   string
		want     []bool
	}{
		{"one_hit", []string{"buy now", "hello"}, "buy", []bool{true, false}},
		{"no_hits", []string{"hello"}, "buy", []bool{false}},
		{"empty_message", []string{""}, "buy", []bool{false}},
		{"all_hits", []string{"buy", "rebuy"}, "buy", []bool{true, true}},
		{"empty", []string{}, "buy", []bool{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Flagged(tc.messages, tc.banned); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Flagged(%v) = %v, want %v", tc.messages, got, tc.want)
			}
		})
	}
}
