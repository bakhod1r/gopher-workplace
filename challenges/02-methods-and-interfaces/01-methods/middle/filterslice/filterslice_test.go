package filterslice

import (
	"reflect"
	"testing"
)

func TestFilterEvens(t *testing.T) {
	cases := []struct {
		name string
		list IntList
		want IntList
	}{
		{"mixed", IntList{1, 2, 3, 4}, IntList{2, 4}},
		{"all_even", IntList{2, 4, 6}, IntList{2, 4, 6}},
		{"all_odd", IntList{1, 3, 5}, IntList{}},
		{"empty", IntList{}, IntList{}},
		{"nil", nil, IntList{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.FilterEvens()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FilterEvens() = %v, want %v", got, tc.want)
			}
		})
	}

	// Verify it returns a new slice, not mutated in place.
	t.Run("immutable", func(t *testing.T) {
		orig := IntList{1, 2, 3}
		got := orig.FilterEvens()
		if len(orig) != 3 {
			t.Errorf("FilterEvens mutated original length: %v", orig)
		}
		if len(got) != 1 || got[0] != 2 {
			t.Errorf("unexpected got: %v", got)
		}
	})
}
