package applyslice

import (
	"reflect"
	"testing"
)

func TestApply(t *testing.T) {
	cases := []struct {
		name   string
		list   IntList
		factor int
		want   IntList
	}{
		{"double", IntList{1, 2, 3}, 2, IntList{2, 4, 6}},
		{"zero", IntList{5, 10}, 0, IntList{0, 0}},
		{"empty", IntList{}, 3, IntList{}},
		{"nil", nil, 5, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Apply(tc.factor)
			if !reflect.DeepEqual(tc.list, tc.want) {
				t.Errorf("Apply(%d) = %v, want %v", tc.factor, tc.list, tc.want)
			}
		})
	}
}
