package emailnormalizer

import (
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		name  string
		addrs []string
		want  []string
	}{
		{"padded_and_upper", []string{" A@X.io "}, []string{"a@x.io"}},
		{"blank", []string{"   "}, []string{""}},
		{"empty", []string{}, []string{}},
		{"tabs_and_newline", []string{"\tB@Y.IO\n"}, []string{"b@y.io"}},
		{"batch", []string{"a@x.io", " C@Z.IO"}, []string{"a@x.io", "c@z.io"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Normalize(tc.addrs); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Normalize(%v) = %v, want %v", tc.addrs, got, tc.want)
			}
		})
	}
}
