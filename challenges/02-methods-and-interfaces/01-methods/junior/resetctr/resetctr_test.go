package resetctr

import "testing"

func TestReset(t *testing.T) {
	cases := []struct {
		name  string
		start int
	}{
		{"from_42", 42},
		{"from_zero", 0},
		{"from_negative", -5},
		{"from_large", 1000000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := Counter{N: tc.start}
			c.Reset()
			if got := c.Value(); got != 0 {
				t.Errorf("Counter{%d}.Reset() => Value() = %d, want 0",
					tc.start, got)
			}
		})
	}
}
