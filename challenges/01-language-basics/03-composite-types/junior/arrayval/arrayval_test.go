package arrayval

import "testing"

func TestSetFirst(t *testing.T) {
	cases := []struct {
		name string
		in   [3]int
		v    int
		want [3]int
	}{
		{"basic", [3]int{1, 2, 3}, 9, [3]int{9, 2, 3}},
		{"zeros", [3]int{0, 0, 0}, 5, [3]int{5, 0, 0}},
		{"same value", [3]int{7, 8, 9}, 7, [3]int{7, 8, 9}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SetFirst(tc.in, tc.v); got != tc.want {
				t.Errorf("SetFirst(%v, %d) = %v, want %v", tc.in, tc.v, got, tc.want)
			}
		})
	}
}

func TestSetFirstDoesNotMutateCaller(t *testing.T) {
	in := [3]int{1, 2, 3}
	_ = SetFirst(in, 99)
	if in != [3]int{1, 2, 3} {
		t.Errorf("caller's array was mutated: got %v, want [1 2 3]", in)
	}
}
