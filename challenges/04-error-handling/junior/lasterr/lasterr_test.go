package lasterr

import "testing"

func TestLastError(t *testing.T) {
	cases := []struct {
		name string
		errs []error
		want error
	}{
		{"nil_slice", nil, nil},
		{"empty", []error{}, nil},
		{"all_nil", []error{nil, nil}, nil},
		{"trailing", []error{ErrA, nil, ErrB}, ErrB},
		{"leading_only", []error{ErrA, nil}, ErrA},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LastError(tc.errs); got != tc.want {
				t.Errorf("LastError(%v) = %v, want %v", tc.errs, got, tc.want)
			}
		})
	}
}
