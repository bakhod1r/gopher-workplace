package applyfn

import (
	"reflect"
	"testing"
)

func TestApplyAll(t *testing.T) {
	cases := []struct {
		name   string
		factor int
		nums   []int
		want   []int
	}{
		{"double", 2, []int{1, 2, 3}, []int{2, 4, 6}},
		{"triple", 3, []int{0, 5}, []int{0, 15}},
		{"empty", 10, []int{}, []int{}},
		{"nil", 10, nil, []int{}},
		{"negative", -1, []int{1, -2, 3}, []int{-1, 2, -3}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tr := Transformer{Factor: tc.factor}
			got := ApplyAll(tr.Transform, tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ApplyAll(×%d, %v) = %v, want %v",
					tc.factor, tc.nums, got, tc.want)
			}
		})
	}
}
