package shardsplit

import (
	"reflect"
	"testing"
)

func TestSplitByShard(t *testing.T) {
	cases := []struct {
		name       string
		userIDs    []int
		wantShard0 []int
		wantShard1 []int
	}{
		{"mixed", []int{1, 2, 3, 4}, []int{2, 4}, []int{1, 3}},
		{"only_even", []int{2}, []int{2}, []int{}},
		{"no_recipients", nil, []int{}, []int{}},
		{"only_odd", []int{1, 3, 5}, []int{}, []int{1, 3, 5}},
		{"negatives_and_zero", []int{-2, -1, 0}, []int{-2, 0}, []int{-1}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got0, got1 := SplitByShard(tc.userIDs)
			if !reflect.DeepEqual(got0, tc.wantShard0) || !reflect.DeepEqual(got1, tc.wantShard1) {
				t.Errorf("SplitByShard(%v) = %#v, %#v, want %#v, %#v",
					tc.userIDs, got0, got1, tc.wantShard0, tc.wantShard1)
			}
		})
	}
}
