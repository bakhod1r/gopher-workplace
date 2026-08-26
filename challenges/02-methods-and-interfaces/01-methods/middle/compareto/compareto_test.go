package compareto

import "testing"

func TestCompare(t *testing.T) {
	cases := []struct {
		name string
		a, b Version
		want int
	}{
		{"less_major", Version{1, 0}, Version{2, 0}, -1},
		{"equal", Version{1, 2}, Version{1, 2}, 0},
		{"greater_major", Version{2, 0}, Version{1, 9}, 1},
		{"less_minor", Version{1, 0}, Version{1, 1}, -1},
		{"greater_minor", Version{1, 5}, Version{1, 3}, 1},
		{"zero", Version{0, 0}, Version{0, 0}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.a.Compare(tc.b); got != tc.want {
				t.Errorf("Version{%d,%d}.Compare({%d,%d}) = %d, want %d",
					tc.a.Major, tc.a.Minor, tc.b.Major, tc.b.Minor, got, tc.want)
			}
		})
	}
}
