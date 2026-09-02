package clampgen

import "testing"

func TestClamp(t *testing.T) {
	cases := []struct {
		name      string
		v, lo, hi int
		want      int
	}{
		{"above", 5, 0, 3, 3},
		{"below", -1, 0, 3, 0},
		{"inside", 2, 0, 3, 2},
		{"at_low", 0, 0, 3, 0},
		{"at_high", 3, 0, 3, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Clamp(tc.v, tc.lo, tc.hi); got != tc.want {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tc.v, tc.lo, tc.hi, got, tc.want)
			}
		})
	}
	if got := Clamp(2.5, 0.0, 1.0); got != 1.0 {
		t.Errorf("Clamp(2.5, 0, 1) = %v, want 1", got)
	}
	if got := Clamp("m", "a", "f"); got != "f" {
		t.Errorf(`Clamp("m", "a", "f") = %q, want "f"`, got)
	}
}
