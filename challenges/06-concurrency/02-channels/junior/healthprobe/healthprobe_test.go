package healthprobe

import (
	"reflect"
	"testing"
)

func TestProbe(t *testing.T) {
	cases := []struct {
		name   string
		rounds int
		want   []string
	}{
		{"one", 1, []string{"probe", "ack"}},
		{"two", 2, []string{"probe", "ack", "probe", "ack"}},
		{"zero", 0, []string{}},
		{"negative", -1, []string{}},
		{"three", 3, []string{"probe", "ack", "probe", "ack", "probe", "ack"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Probe(tc.rounds)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Probe(%d) = %#v, want %#v", tc.rounds, got, tc.want)
			}
		})
	}
}
