package issortedgen

import "testing"

func TestIsSorted(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		want bool
	}{
		{"ascending", []int{1, 2, 3}, true},
		{"duplicates_ok", []int{1, 2, 2}, true},
		{"descending", []int{2, 1}, false},
		{"single", []int{5}, true},
		{"empty", []int{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSorted(tc.s); got != tc.want {
				t.Errorf("IsSorted(%v) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
	if !IsSorted([]string{"a", "b"}) {
		t.Error(`IsSorted([]string{"a", "b"}) = false, want true`)
	}
}
