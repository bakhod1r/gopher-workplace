package healthfanout

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

var errDown = errors.New("connection refused")

// prober fails for every service listed in down.
func prober(down []string, calls *int64) func(string) error {
	set := map[string]bool{}
	for _, s := range down {
		set[s] = true
	}
	return func(svc string) error {
		atomic.AddInt64(calls, 1)
		if set[svc] {
			return errDown
		}
		return nil
	}
}

func TestUnhealthyServices(t *testing.T) {
	cases := []struct {
		name     string
		services []string
		down     []string
		want     []string
	}{
		{"all_healthy", []string{"api", "db", "queue"}, nil, []string{}},
		{"one_down", []string{"api", "db"}, []string{"db"}, []string{"db"}},
		{"all_down", []string{"web", "api"}, []string{"web", "api"}, []string{"api", "web"}},
		{"result_is_sorted", []string{"zeta", "alpha", "mid"}, []string{"zeta", "alpha"}, []string{"alpha", "zeta"}},
		{"single_service", []string{"api"}, []string{"api"}, []string{"api"}},
		{"no_services", nil, nil, []string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			got := UnhealthyServices(tc.services, prober(tc.down, &calls))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("UnhealthyServices(%v) = %v, want %v", tc.services, got, tc.want)
			}
			if int(calls) != len(tc.services) {
				t.Errorf("probe called %d times, want %d", calls, len(tc.services))
			}
		})
	}
}
