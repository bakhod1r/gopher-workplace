package firstfail

import (
	"errors"
	"testing"
)

func TestFirstFail(t *testing.T) {
	cases := []struct {
		name    string
		errs    []error
		want    int
		wantErr error
	}{
		{"second", []error{nil, ErrStep}, 1, nil},
		{"first", []error{ErrStep, ErrStep}, 0, nil},
		{"none", []error{nil, nil}, -1, ErrNoFailure},
		{"empty", []error{}, -1, ErrNoFailure},
		{"nil_slice", nil, -1, ErrNoFailure},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FirstFail(tc.errs)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("FirstFail(%v) err = %v, want %v", tc.errs, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("FirstFail(%v) = %d, want %d", tc.errs, got, tc.want)
			}
		})
	}
}
