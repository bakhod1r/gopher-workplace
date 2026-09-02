package popstack

import (
	"errors"
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	cases := []struct {
		name    string
		s       []int
		wantS   []int
		wantV   int
		wantErr error
	}{
		{"three", []int{1, 2, 3}, []int{1, 2}, 3, nil},
		{"one", []int{7}, []int{}, 7, nil},
		{"empty", []int{}, nil, 0, ErrEmpty},
		{"nil", nil, nil, 0, ErrEmpty},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotS, gotV, err := Pop(tc.s)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Pop(%v) err = %v, want %v", tc.s, err, tc.wantErr)
			}
			if gotV != tc.wantV {
				t.Errorf("Pop(%v) value = %d, want %d", tc.s, gotV, tc.wantV)
			}
			if err == nil && !reflect.DeepEqual(gotS, tc.wantS) {
				t.Errorf("Pop(%v) rest = %v, want %v", tc.s, gotS, tc.wantS)
			}
		})
	}
}
