package backupuploader

import (
	"reflect"
	"testing"
)

func TestPartChecksums(t *testing.T) {
	cases := []struct {
		name  string
		parts []string
		want  []int
	}{
		{"single_bytes", []string{"a", "b"}, []int{97, 98}},
		{"two_bytes", []string{"ab"}, []int{3105}},
		{"empty_part", []string{""}, []int{0}},
		{"empty", []string{}, []int{}},
		{"repeat", []string{"aa", "a"}, []int{3104, 97}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PartChecksums(tc.parts); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PartChecksums(%v) = %v, want %v", tc.parts, got, tc.want)
			}
		})
	}
}
