package containsgen

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		v    int
		want bool
	}{
		{"present_middle", []int{1, 2, 3}, 2, true},
		{"present_first", []int{1, 2, 3}, 1, true},
		{"present_last", []int{1, 2, 3}, 3, true},
		{"absent", []int{1, 2, 3}, 9, false},
		{"empty", []int{}, 1, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.s, tc.v); got != tc.want {
				t.Errorf("Contains(%v, %v) = %v, want %v", tc.s, tc.v, got, tc.want)
			}
		})
	}
	if Contains([]string{"a"}, "b") {
		t.Error(`Contains([]string{"a"}, "b") = true, want false`)
	}
	if !Contains([]string{"a", "b"}, "b") {
		t.Error(`Contains([]string{"a", "b"}, "b") = false, want true`)
	}
}
