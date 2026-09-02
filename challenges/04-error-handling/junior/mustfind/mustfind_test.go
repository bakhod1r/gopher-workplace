package mustfind

import (
	"errors"
	"testing"
)

func TestFind(t *testing.T) {
	cases := []struct {
		name    string
		s       []int
		target  int
		want    int
		wantErr error
	}{
		{"middle", []int{4, 7, 9}, 7, 1, nil},
		{"first", []int{4, 7}, 4, 0, nil},
		{"duplicate", []int{5, 5, 5}, 5, 0, nil},
		{"absent", []int{4, 7}, 5, -1, ErrNotFound},
		{"empty", []int{}, 1, -1, ErrNotFound},
		{"nil_slice", nil, 1, -1, ErrNotFound},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Find(tc.s, tc.target)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Find(%v, %d) err = %v, want %v", tc.s, tc.target, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Find(%v, %d) = %d, want %d", tc.s, tc.target, got, tc.want)
			}
		})
	}
}
