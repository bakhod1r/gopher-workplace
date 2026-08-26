package increment

import "testing"

func TestInc(t *testing.T) {
	cases := []struct {
		name  string
		start int
		calls int
		want  int
	}{
		{"from_zero", 0, 1, 1},
		{"from_five", 5, 1, 6},
		{"triple", 0, 3, 3},
		{"ten", 10, 5, 15},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := Counter{N: tc.start}
			for i := 0; i < tc.calls; i++ {
				c.Inc()
			}
			if got := c.Value(); got != tc.want {
				t.Errorf("Counter{%d} after %d Inc() = %d, want %d",
					tc.start, tc.calls, got, tc.want)
			}
		})
	}
}
