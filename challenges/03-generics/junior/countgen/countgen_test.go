package countgen

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		v    int
		want int
	}{
		{"two_matches", []int{1, 2, 1}, 1, 2},
		{"no_match", []int{1, 2}, 9, 0},
		{"all_match", []int{4, 4, 4}, 4, 3},
		{"empty", []int{}, 1, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.s, tc.v); got != tc.want {
				t.Errorf("Count(%v, %v) = %v, want %v", tc.s, tc.v, got, tc.want)
			}
		})
	}
	if got := Count([]string{"a", "a", "b"}, "a"); got != 2 {
		t.Errorf(`Count([]string{"a", "a", "b"}, "a") = %v, want 2`, got)
	}
}
