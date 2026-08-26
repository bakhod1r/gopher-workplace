package methodexpr

import "testing"

func TestCallExpr(t *testing.T) {
	cases := []struct {
		name string
		base int
		n    int
		want int
	}{
		{"10+5", 10, 5, 15},
		{"0+0", 0, 0, 0},
		{"neg", -3, 7, 4},
		{"large", 100, 200, 300},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CallExpr(Adder.Add, Adder{Base: tc.base}, tc.n)
			if got != tc.want {
				t.Errorf("CallExpr(Adder.Add, Adder{%d}, %d) = %d, want %d",
					tc.base, tc.n, got, tc.want)
			}
		})
	}
}
