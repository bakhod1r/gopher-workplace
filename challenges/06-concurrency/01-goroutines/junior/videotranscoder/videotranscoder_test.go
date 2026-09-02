package videotranscoder

import (
	"reflect"
	"testing"
)

func TestTargetBitrates(t *testing.T) {
	cases := []struct {
		name      string
		bitrates  []int
		factorPct int
		want      []int
	}{
		{"half", []int{4000, 2000}, 50, []int{2000, 1000}},
		{"unchanged", []int{4000}, 100, []int{4000}},
		{"zero_factor", []int{4000, 1}, 0, []int{0, 0}},
		{"negative_factor", []int{4000}, -20, []int{0}},
		{"empty", []int{}, 50, []int{}},
		{"quarter", []int{800, 400, 200}, 25, []int{200, 100, 50}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TargetBitrates(tc.bitrates, tc.factorPct); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TargetBitrates(%v) = %v, want %v", tc.bitrates, got, tc.want)
			}
		})
	}
}
