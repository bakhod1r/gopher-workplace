package keygenerator

import (
	"reflect"
	"testing"
)

func TestPrimeCandidates(t *testing.T) {
	cases := []struct {
		name       string
		candidates []int
		want       []bool
	}{
		{"small_primes", []int{2, 4, 7}, []bool{true, false, true}},
		{"below_two", []int{1, 0, -5}, []bool{false, false, false}},
		{"perfect_squares", []int{9, 25}, []bool{false, false}},
		{"larger", []int{97, 91}, []bool{true, false}},
		{"empty", []int{}, []bool{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PrimeCandidates(tc.candidates); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PrimeCandidates(%v) = %v, want %v", tc.candidates, got, tc.want)
			}
		})
	}
}
