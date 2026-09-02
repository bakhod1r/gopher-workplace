package elemat

import (
	"errors"
	"testing"
)

func TestAt(t *testing.T) {
	s := []int{1, 2, 3}

	cases := []struct {
		name    string
		s       []int
		i       int
		want    int
		wantErr error
	}{
		{"first", s, 0, 1, nil},
		{"last", s, 2, 3, nil},
		{"past_end", s, 3, 0, ErrOutOfRange},
		{"negative", s, -1, 0, ErrOutOfRange},
		{"nil_slice", nil, 0, 0, ErrOutOfRange},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := At(tc.s, tc.i)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("At(%v, %d) err = %v, want %v", tc.s, tc.i, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("At(%v, %d) = %d, want %d", tc.s, tc.i, got, tc.want)
			}
		})
	}
}
