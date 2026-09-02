package startupsequence

import (
	"reflect"
	"testing"
)

func TestRunChecks(t *testing.T) {
	configOK := func() int { return 0 }
	diskFull := func() int { return 28 }
	noNetwork := func() int { return 101 }

	cases := []struct {
		name   string
		checks []func() int
		want   []int
	}{
		{"ok_and_failure", []func() int{configOK, diskFull}, []int{0, 28}},
		{"single_ok", []func() int{configOK}, []int{0}},
		{"empty", []func() int{}, []int{}},
		{"order_preserved", []func() int{noNetwork, configOK, diskFull}, []int{101, 0, 28}},
		{"all_failing", []func() int{diskFull, noNetwork}, []int{28, 101}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RunChecks(tc.checks); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RunChecks(%v) = %v, want %v", "checks", got, tc.want)
			}
		})
	}
}
