package absval

import "testing"

func TestAbs(t *testing.T) {
	cases := []struct {
		name string
		n    MyInt
		want MyInt
	}{
		{"negative", -5, 5},
		{"positive", 3, 3},
		{"zero", 0, 0},
		{"minus_one", -1, 1},
		{"large_neg", -1000, 1000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.n.Abs(); got != tc.want {
				t.Errorf("MyInt(%d).Abs() = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
