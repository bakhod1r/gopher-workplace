package counterrs

import "testing"

func TestCountErrors(t *testing.T) {
	cases := []struct {
		name string
		errs []error
		want int
	}{
		{"nil_slice", nil, 0},
		{"empty", []error{}, 0},
		{"none", []error{nil, nil}, 0},
		{"two", []error{nil, ErrX, ErrX}, 2},
		{"all", []error{ErrX, ErrX, ErrX}, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountErrors(tc.errs); got != tc.want {
				t.Errorf("CountErrors(%v) = %d, want %d", tc.errs, got, tc.want)
			}
		})
	}
}
