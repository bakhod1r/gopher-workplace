package healthcheck

import (
	"strings"
	"testing"
)

func TestCountHealthy(t *testing.T) {
	isOK := func(host string) bool { return strings.HasPrefix(host, "ok-") }

	cases := []struct {
		name  string
		hosts []string
		want  int
	}{
		{"one_of_two", []string{"ok-a", "bad"}, 1},
		{"all_healthy", []string{"ok-a", "ok-b"}, 2},
		{"none_healthy", []string{"bad-a", "bad-b"}, 0},
		{"fleet_of_five", []string{"ok-a", "ok-b", "bad", "ok-c", "bad-2"}, 3},
		{"empty_fleet", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountHealthy(tc.hosts, isOK); got != tc.want {
				t.Errorf("CountHealthy(%v) = %d, want %d", tc.hosts, got, tc.want)
			}
		})
	}
}
