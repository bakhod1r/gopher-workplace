package healthpings

import "testing"

func TestHealthPings(t *testing.T) {
	cases := []struct {
		name     string
		endpoint string
		count    int
		want     []string
	}{
		{"two_probes", "api", 2, []string{"api/health", "api/health"}},
		{"one_probe", "db", 1, []string{"db/health"}},
		{"three_probes", "cache", 3, []string{"cache/health", "cache/health", "cache/health"}},
		{"zero_probes", "api", 0, nil},
		{"negative_count", "api", -1, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for p := range HealthPings(tc.endpoint, tc.count) {
				got = append(got, p)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("HealthPings(%q, %d) = %v, want %v", tc.endpoint, tc.count, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("HealthPings(%q, %d) = %v, want %v", tc.endpoint, tc.count, got, tc.want)
				}
			}
		})
	}
}
