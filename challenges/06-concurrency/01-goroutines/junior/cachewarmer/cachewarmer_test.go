package cachewarmer

import (
	"reflect"
	"testing"
)

func TestWarmAll(t *testing.T) {
	loader := func(key string) int {
		if key == "missing" {
			return -1
		}
		return len(key) * 100
	}

	cases := []struct {
		name     string
		keys     []string
		load     func(string) int
		capBytes int
		want     []int
	}{
		{"under_cap", []string{"a", "bb"}, loader, 1000, []int{100, 200}},
		{"clamped", []string{"huge"}, loader, 150, []int{150}},
		{"empty", []string{}, loader, 100, []int{}},
		{"miss_is_zero", []string{"missing"}, loader, 500, []int{0}},
		{"mixed", []string{"a", "missing", "bbb"}, loader, 250, []int{100, 0, 250}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WarmAll(tc.keys, tc.load, tc.capBytes); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("WarmAll(%v) = %v, want %v", tc.keys, got, tc.want)
			}
		})
	}
}
