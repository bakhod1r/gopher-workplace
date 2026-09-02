package firsterr

import "testing"

func TestFirstError(t *testing.T) {
	cases := []struct {
		name string
		errs []error
		want error
	}{
		{"nil_slice", nil, nil},
		{"empty", []error{}, nil},
		{"all_nil", []error{nil, nil, nil}, nil},
		{"second", []error{nil, ErrB, ErrC}, ErrB},
		{"first", []error{ErrA, ErrB}, ErrA},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstError(tc.errs); got != tc.want {
				t.Errorf("FirstError(%v) = %v, want %v", tc.errs, got, tc.want)
			}
		})
	}
}
