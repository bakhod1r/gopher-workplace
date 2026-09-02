package dnsresolver

import (
	"reflect"
	"testing"
)

func TestResolveAll(t *testing.T) {
	lookup := func(host string) string {
		switch host {
		case "a.io":
			return "10.0.0.1"
		case "b.io":
			return "10.0.0.2"
		}
		return ""
	}

	cases := []struct {
		name    string
		hosts   []string
		resolve func(string) string
		want    []string
	}{
		{"both_known", []string{"a.io", "b.io"}, lookup, []string{"10.0.0.1", "10.0.0.2"}},
		{"unknown_host", []string{"ghost.io"}, lookup, []string{"0.0.0.0"}},
		{"empty", []string{}, lookup, []string{}},
		{"mixed", []string{"a.io", "ghost.io"}, lookup, []string{"10.0.0.1", "0.0.0.0"}},
		{"duplicates", []string{"b.io", "b.io"}, lookup, []string{"10.0.0.2", "10.0.0.2"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ResolveAll(tc.hosts, tc.resolve); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ResolveAll(%v) = %v, want %v", tc.hosts, got, tc.want)
			}
		})
	}
}
