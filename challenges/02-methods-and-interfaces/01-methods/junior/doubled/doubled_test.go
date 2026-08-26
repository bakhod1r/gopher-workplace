package doubled

import "testing"

func TestDouble(t *testing.T) {
	cases := []struct {
		name string
		f    MyFloat
		want MyFloat
	}{
		{"positive", 3.5, 7.0},
		{"zero", 0, 0},
		{"negative", -2.5, -5.0},
		{"one", 1, 2},
		{"small", 0.1, 0.2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.f.Double(); got != tc.want {
				t.Errorf("MyFloat(%g).Double() = %g, want %g",
					float64(tc.f), float64(got), float64(tc.want))
			}
		})
	}
}
