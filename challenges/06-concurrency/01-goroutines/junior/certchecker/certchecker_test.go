package certchecker

import (
	"reflect"
	"testing"
)

func TestExpiringSoon(t *testing.T) {
	cases := []struct {
		name     string
		expiries []int
		today    int
		window   int
		want     []bool
	}{
		{"one_soon_one_far", []int{100, 400}, 90, 30, []bool{true, false}},
		{"already_expired", []int{50}, 90, 30, []bool{true}},
		{"exactly_on_window", []int{120}, 90, 30, []bool{true}},
		{"just_outside", []int{121}, 90, 30, []bool{false}},
		{"empty", []int{}, 90, 30, []bool{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ExpiringSoon(tc.expiries, tc.today, tc.window); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ExpiringSoon(%v) = %v, want %v", tc.expiries, got, tc.want)
			}
		})
	}
}
