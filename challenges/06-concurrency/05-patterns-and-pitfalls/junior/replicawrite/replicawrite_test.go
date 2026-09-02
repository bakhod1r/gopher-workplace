package replicawrite

import (
	"strings"
	"testing"
)

func TestFirstReplicaAck(t *testing.T) {
	write := func(replica string) bool { return strings.HasPrefix(replica, "ok-") }

	cases := []struct {
		name     string
		replicas []string
		want     bool
	}{
		{"one_of_two_acks", []string{"ok-1", "bad-2"}, true},
		{"all_ack", []string{"ok-1", "ok-2"}, true},
		{"none_ack", []string{"bad-1", "bad-2"}, false},
		{"single_failure", []string{"bad-1"}, false},
		{"no_replicas", nil, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstReplicaAck(tc.replicas, write); got != tc.want {
				t.Errorf("FirstReplicaAck(%v) = %v, want %v", tc.replicas, got, tc.want)
			}
		})
	}
}
