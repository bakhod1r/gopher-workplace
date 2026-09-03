package regionreplicator

import (
	"errors"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
)

var errZoneUnreachable = errors.New("zone unreachable")

// replicator fails for any zone whose name contains "bad".
func replicator(calls *int64) func(string, string) error {
	return func(region, zone string) error {
		atomic.AddInt64(calls, 1)
		if strings.Contains(zone, "bad") {
			return errZoneUnreachable
		}
		return nil
	}
}

func zoneCount(regions []Region) int {
	n := 0
	for _, r := range regions {
		n += len(r.Zones)
	}
	return n
}

func TestReplicateAll(t *testing.T) {
	cases := []struct {
		name    string
		regions []Region
		want    []string
	}{
		{
			"all_zones_replicate",
			[]Region{{"eu", []string{"a", "b"}}, {"us", []string{"c"}}},
			[]string{},
		},
		{
			"one_zone_down",
			[]Region{{"eu", []string{"a", "bad-b"}}, {"us", []string{"c"}}},
			[]string{"eu/bad-b"},
		},
		{
			"failures_sorted_across_regions",
			[]Region{{"us", []string{"bad-1"}}, {"ap", []string{"bad-2"}}, {"eu", []string{"ok"}}},
			[]string{"ap/bad-2", "us/bad-1"},
		},
		{
			"whole_region_down",
			[]Region{{"eu", []string{"bad-a", "bad-b"}}},
			[]string{"eu/bad-a", "eu/bad-b"},
		},
		{
			"region_with_no_zones",
			[]Region{{"eu", nil}, {"us", []string{"c"}}},
			[]string{},
		},
		{"no_regions", nil, []string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			got := ReplicateAll(tc.regions, replicator(&calls))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ReplicateAll() = %v, want %v", got, tc.want)
			}
			if int(calls) != zoneCount(tc.regions) {
				t.Errorf("replicate called %d times, want %d", calls, zoneCount(tc.regions))
			}
		})
	}
}
