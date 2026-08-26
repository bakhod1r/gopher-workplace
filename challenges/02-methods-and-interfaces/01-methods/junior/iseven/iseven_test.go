package iseven

import "testing"

func TestIsEven(t *testing.T) {
	cases := []struct {
		name string
		n    MyInt
		want bool
	}{
		{"four", 4, true},
		{"three", 3, false},
		{"zero", 0, true},
		{"negative_even", -2, true},
		{"negative_odd", -7, false},
		{"one", 1, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.n.IsEven(); got != tc.want {
				t.Errorf("MyInt(%d).IsEven() = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
