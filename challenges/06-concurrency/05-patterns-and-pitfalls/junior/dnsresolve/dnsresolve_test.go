package dnsresolve

import (
	"reflect"
	"testing"
)

func TestResolveAll(t *testing.T) {
	lookup := func(host string) string { return "10.0.0." + host }

	cases := []struct {
		name  string
		hosts []string
		want  map[string]string
	}{
		{"single_host", []string{"1"}, map[string]string{"1": "10.0.0.1"}},
		{"two_hosts", []string{"1", "2"}, map[string]string{"1": "10.0.0.1", "2": "10.0.0.2"}},
		{"duplicate_hosts", []string{"1", "1"}, map[string]string{"1": "10.0.0.1"}},
		{"five_hosts", []string{"1", "2", "3", "4", "5"}, map[string]string{"1": "10.0.0.1", "2": "10.0.0.2", "3": "10.0.0.3", "4": "10.0.0.4", "5": "10.0.0.5"}},
		{"no_hosts", nil, map[string]string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ResolveAll(tc.hosts, lookup)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ResolveAll(%v) = %v, want %v", tc.hosts, got, tc.want)
			}
		})
	}
}
