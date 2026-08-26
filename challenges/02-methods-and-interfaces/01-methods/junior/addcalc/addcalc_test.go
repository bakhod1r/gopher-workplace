package addcalc

import "testing"

func TestAdd(t *testing.T) {
	c := Calculator{}
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"zero_sum", -1, 1, 0},
		{"both_zero", 0, 0, 0},
		{"negatives", -5, -3, -8},
		{"large", 1000000, 2000000, 3000000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := c.Add(tc.a, tc.b); got != tc.want {
				t.Errorf("Calculator.Add(%d, %d) = %d, want %d",
					tc.a, tc.b, got, tc.want)
			}
		})
	}
}
