package ratelimiter

import (
	"reflect"
	"testing"
)

func TestRemainingQuota(t *testing.T) {
	cases := []struct {
		name  string
		used  []int
		limit int
		want  []int
	}{
		{"within_limit", []int{10, 90}, 100, []int{90, 10}},
		{"over_limit", []int{150}, 100, []int{0}},
		{"exactly_spent", []int{100}, 100, []int{0}},
		{"nothing_used", []int{0, 0}, 50, []int{50, 50}},
		{"empty", []int{}, 100, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RemainingQuota(tc.used, tc.limit); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RemainingQuota(%v) = %v, want %v", tc.used, got, tc.want)
			}
		})
	}
}
