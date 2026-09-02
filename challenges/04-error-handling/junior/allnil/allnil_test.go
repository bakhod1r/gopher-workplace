package allnil

import "testing"

func TestAllNil(t *testing.T) {
	cases := []struct {
		name string
		errs []error
		want bool
	}{
		{"nil_slice", nil, true},
		{"empty", []error{}, true},
		{"all_nil", []error{nil, nil, nil}, true},
		{"one_failure", []error{nil, ErrCheck}, false},
		{"first_failure", []error{ErrCheck, nil}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := AllNil(tc.errs); got != tc.want {
				t.Errorf("AllNil(%v) = %v, want %v", tc.errs, got, tc.want)
			}
		})
	}
}
