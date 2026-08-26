package mergeobj

import "testing"

func TestMerge(t *testing.T) {
	cases := []struct {
		name  string
		base  Config
		other Config
		want  Config
	}{
		{"port_only", Config{"localhost", 8080, false},
			Config{Port: 9090}, Config{"localhost", 9090, false}},
		{"debug_only", Config{"host", 80, false},
			Config{Debug: true}, Config{"host", 80, true}},
		{"host_and_port", Config{"old", 80, false},
			Config{Host: "new", Port: 443}, Config{"new", 443, false}},
		{"no_change", Config{"h", 80, true},
			Config{}, Config{"h", 80, true}},
		{"all", Config{"a", 1, false},
			Config{"b", 2, true}, Config{"b", 2, true}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := tc.base
			c.Merge(tc.other)
			if c != tc.want {
				t.Errorf("Merge(%+v) => %+v, want %+v", tc.other, c, tc.want)
			}
		})
	}
}
