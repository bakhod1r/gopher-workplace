package sensorfailover

import (
	"reflect"
	"sort"
	"testing"
)

func streamOf(names ...string) <-chan Reading {
	ch := make(chan Reading, len(names))
	for i, n := range names {
		ch <- Reading{Sensor: n, Celsius: float64(i)}
	}
	close(ch)
	return ch
}

func sensors(rs []Reading) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.Sensor)
	}
	sort.Strings(out)
	return out
}

func TestMergeSensorStreams(t *testing.T) {
	cases := []struct {
		name            string
		primary, backup []string
		nilPrimary      bool
		nilBackup       bool
		want            []string
	}{
		{"both_deliver", []string{"t1", "t2"}, []string{"t3"}, false, false, []string{"t1", "t2", "t3"}},
		{"primary_closes_empty", nil, []string{"t3", "t4"}, false, false, []string{"t3", "t4"}},
		{"backup_closes_empty", []string{"t1"}, nil, false, false, []string{"t1"}},
		{"both_empty", nil, nil, false, false, []string{}},
		{"backup_never_wired", []string{"t1", "t2"}, nil, false, true, []string{"t1", "t2"}},
		{"primary_never_wired", nil, []string{"t9"}, true, false, []string{"t9"}},
		{"both_nil", nil, nil, true, true, []string{}},
		{"long_run", []string{"a", "b", "c", "d"}, []string{"e", "f"}, false, false,
			[]string{"a", "b", "c", "d", "e", "f"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p, b <-chan Reading
			if !tc.nilPrimary {
				p = streamOf(tc.primary...)
			}
			if !tc.nilBackup {
				b = streamOf(tc.backup...)
			}
			got := MergeSensorStreams(p, b)
			if !reflect.DeepEqual(sensors(got), tc.want) {
				t.Errorf("MergeSensorStreams() sensors = %#v, want %#v", sensors(got), tc.want)
			}
		})
	}
}
