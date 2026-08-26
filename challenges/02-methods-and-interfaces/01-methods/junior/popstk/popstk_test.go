package popstk

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	cases := []struct {
		name     string
		items    []int
		wantVal  int
		wantOK   bool
		wantLeft []int
	}{
		{"pop_last", []int{1, 2, 3}, 3, true, []int{1, 2}},
		{"pop_single", []int{42}, 42, true, []int{}},
		{"empty", nil, 0, false, nil},
		{"empty_slice", []int{}, 0, false, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stack{Items: tc.items}
			v, ok := s.Pop()
			if v != tc.wantVal || ok != tc.wantOK {
				t.Errorf("Pop() = (%d, %v), want (%d, %v)",
					v, ok, tc.wantVal, tc.wantOK)
			}
			if !reflect.DeepEqual(s.Items, tc.wantLeft) {
				t.Errorf("after Pop(): Items = %v, want %v", s.Items, tc.wantLeft)
			}
		})
	}
}
