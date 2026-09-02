package aspectratio

import (
	"reflect"
	"testing"
)

func TestDivisors(t *testing.T) {
	cases := []struct {
		name  string
		sizes [][2]int
		want  []int
	}{
		{"full_hd", [][2]int{{1920, 1080}}, []int{120}},
		{"zero_width", [][2]int{{0, 720}}, []int{720}},
		{"coprime", [][2]int{{7, 13}}, []int{1}},
		{"square", [][2]int{{800, 800}}, []int{800}},
		{"empty", [][2]int{}, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Divisors(tc.sizes); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Divisors(%v) = %v, want %v", tc.sizes, got, tc.want)
			}
		})
	}
}
