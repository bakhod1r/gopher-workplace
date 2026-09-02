package indexofgen

import "testing"

func TestIndexOf(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		v    int
		want int
	}{
		{"first_of_duplicates", []int{5, 7, 7}, 7, 1},
		{"index_zero", []int{5, 7}, 5, 0},
		{"absent", []int{5}, 7, -1},
		{"empty", []int{}, 7, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IndexOf(tc.s, tc.v); got != tc.want {
				t.Errorf("IndexOf(%v, %v) = %v, want %v", tc.s, tc.v, got, tc.want)
			}
		})
	}
	if got := IndexOf([]string{"a", "b"}, "b"); got != 1 {
		t.Errorf(`IndexOf([]string{"a", "b"}, "b") = %v, want 1`, got)
	}
}
