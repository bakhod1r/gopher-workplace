package checkoutgrid

import (
	"reflect"
	"testing"
)

func TestSeatMap(t *testing.T) {
	cases := []struct {
		name  string
		taken [][2]int
		want  func() [7][10]bool
	}{
		{"single", [][2]int{{0, 0}}, func() [7][10]bool { var g [7][10]bool; g[0][0] = true; return g }},
		{"several", [][2]int{{1, 2}, {6, 9}}, func() [7][10]bool { var g [7][10]bool; g[1][2] = true; g[6][9] = true; return g }},
		{"out of range ignored", [][2]int{{7, 0}, {0, 10}, {-1, 0}}, func() [7][10]bool { var g [7][10]bool; return g }},
		{"empty", nil, func() [7][10]bool { var g [7][10]bool; return g }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SeatMap(tc.taken); !reflect.DeepEqual(got, tc.want()) {
				t.Errorf("SeatMap(%v) = %v, want %v", tc.taken, got, tc.want())
			}
		})
	}
}
