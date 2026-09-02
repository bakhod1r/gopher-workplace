package allowlist

import (
	"strconv"
	"sync"
	"testing"
)

func TestAllowlist(t *testing.T) {
	cases := []struct {
		name     string
		existing []string
		add      string
		wantNew  bool
		wantSize int
	}{
		{"first_ip", nil, "10.0.0.1", true, 1},
		{"duplicate", []string{"10.0.0.1"}, "10.0.0.1", false, 1},
		{"second_ip", []string{"10.0.0.1"}, "10.0.0.2", true, 2},
		{"duplicate_of_many", []string{"a", "b", "c"}, "b", false, 3},
		{"new_after_many", []string{"a", "b", "c"}, "d", true, 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := NewAllowlist()
			for _, ip := range tc.existing {
				a.Allow(ip)
			}
			if got := a.Allow(tc.add); got != tc.wantNew {
				t.Errorf("Allow(%q) = %v, want %v", tc.add, got, tc.wantNew)
			}
			if !a.Allowed(tc.add) {
				t.Errorf("Allowed(%q) = false, want true", tc.add)
			}
			if got := a.Size(); got != tc.wantSize {
				t.Errorf("Size() = %d, want %d", got, tc.wantSize)
			}
		})
	}
}

func TestAllowlistConcurrent(t *testing.T) {
	a := NewAllowlist()
	const admins = 16
	const ips = 25
	added := make(chan string, admins*ips)
	var wg sync.WaitGroup
	wg.Add(admins)
	for i := 0; i < admins; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < ips; j++ {
				ip := "10.0.0." + strconv.Itoa(j)
				if a.Allow(ip) {
					added <- ip
				}
				a.Allowed(ip)
			}
		}()
	}
	wg.Wait()
	close(added)

	if got := len(added); got != ips {
		t.Errorf("newly added IPs = %d, want %d", got, ips)
	}
	if got := a.Size(); got != ips {
		t.Errorf("Size() = %d, want %d", got, ips)
	}
}
