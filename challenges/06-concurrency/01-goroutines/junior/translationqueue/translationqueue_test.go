package translationqueue

import (
	"reflect"
	"testing"
)

func TestCharCounts(t *testing.T) {
	cases := []struct {
		name     string
		messages []string
		want     []int
	}{
		{"ascii_and_accent", []string{"go", "añb"}, []int{2, 3}},
		{"cjk", []string{"日本"}, []int{2}},
		{"empty_message", []string{""}, []int{0}},
		{"empty", []string{}, []int{}},
		{"emoji", []string{"a🙂", "abc"}, []int{2, 3}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CharCounts(tc.messages); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CharCounts(%v) = %v, want %v", tc.messages, got, tc.want)
			}
		})
	}
}
