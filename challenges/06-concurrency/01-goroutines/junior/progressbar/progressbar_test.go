package progressbar

import (
	"reflect"
	"testing"
)

func TestBars(t *testing.T) {
	cases := []struct {
		name     string
		percents []int
		width    int
		want     []string
	}{
		{"half_and_full", []int{50, 100}, 10, []string{"#####", "##########"}},
		{"zero", []int{0}, 10, []string{""}},
		{"over_and_under", []int{150, -20}, 10, []string{"##########", ""}},
		{"narrow_bar", []int{25}, 4, []string{"#"}},
		{"empty", []int{}, 10, []string{}},
		{"zero_width", []int{50}, 0, []string{""}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Bars(tc.percents, tc.width); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Bars(%v) = %v, want %v", tc.percents, got, tc.want)
			}
		})
	}
}
