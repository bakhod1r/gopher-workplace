package scale

import "testing"

func TestScale(t *testing.T) {
	cases := []struct {
		name   string
		v      Vector
		factor float64
		want   Vector
	}{
		{"double", Vector{3, 4}, 2, Vector{6, 8}},
		{"zero", Vector{1, -1}, 0, Vector{0, 0}},
		{"half", Vector{10, 6}, 0.5, Vector{5, 3}},
		{"negative", Vector{2, 3}, -1, Vector{-2, -3}},
		{"identity", Vector{7, 8}, 1, Vector{7, 8}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := tc.v
			v.Scale(tc.factor)
			if v != tc.want {
				t.Errorf("Vector{%g,%g}.Scale(%g) => {%g,%g}, want {%g,%g}",
					tc.v.X, tc.v.Y, tc.factor, v.X, v.Y, tc.want.X, tc.want.Y)
			}
		})
	}
}
