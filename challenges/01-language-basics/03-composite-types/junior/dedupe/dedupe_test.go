package dedupe

import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"adjacent dups", []int{1, 1, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"first-appearance order", []int{5, 4, 5, 4}, []int{5, 4}},
		{"all same", []int{7, 7, 7}, []int{7}},
		{"no dups", []int{1, 2, 3}, []int{1, 2, 3}},
		{"nil", nil, []int{}},
		{"empty", []int{}, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Dedupe(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Dedupe(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestDedupeDoesNotMutateInput(t *testing.T) {
	in := []int{1, 1, 2, 3}
	snapshot := []int{1, 1, 2, 3}
	_ = Dedupe(in)
	if !reflect.DeepEqual(in, snapshot) {
		t.Errorf("input was mutated: got %v, want %v", in, snapshot)
	}
}
