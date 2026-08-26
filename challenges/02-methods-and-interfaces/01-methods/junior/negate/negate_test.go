package negate

import "testing"

func TestNegate(t *testing.T) {
	cases := []struct {
		name string
		n    MyInt
		want MyInt
	}{
		{"positive", 5, -5},
		{"negative", -3, 3},
		{"zero", 0, 0},
		{"one", 1, -1},
		{"minus_one", -1, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n := tc.n
			n.Negate()
			if n != tc.want {
				t.Errorf("MyInt(%d).Negate() => %d, want %d", tc.n, n, tc.want)
			}
		})
	}
}
