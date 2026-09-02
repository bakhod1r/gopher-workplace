package endpointhits

import (
	"sync"
	"testing"
)

func TestHitCounter(t *testing.T) {
	cases := []struct {
		name    string
		records []string
		query   string
		want    int
	}{
		{"one_request", []string{"/users"}, "/users", 1},
		{"two_requests", []string{"/users", "/users"}, "/users", 2},
		{"unserved_route", []string{"/users"}, "/orders", 0},
		{"no_traffic", nil, "/users", 0},
		{"mixed_routes", []string{"/users", "/orders", "/users", "/health"}, "/users", 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHitCounter()
			for _, r := range tc.records {
				h.Record(r)
			}
			if got := h.Hits(tc.query); got != tc.want {
				t.Errorf("Hits(%q) = %d, want %d", tc.query, got, tc.want)
			}
		})
	}
}

func TestHitCounterConcurrent(t *testing.T) {
	h := NewHitCounter()
	routes := []string{"/users", "/orders", "/health"}
	const handlers = 12
	const per = 100
	var wg sync.WaitGroup
	wg.Add(handlers)
	for i := 0; i < handlers; i++ {
		go func(i int) {
			defer wg.Done()
			route := routes[i%len(routes)]
			for j := 0; j < per; j++ {
				h.Record(route)
				h.Hits(route)
			}
		}(i)
	}
	wg.Wait()

	total := 0
	for _, r := range routes {
		total += h.Hits(r)
	}
	if want := handlers * per; total != want {
		t.Errorf("total hits = %d, want %d", total, want)
	}
}
