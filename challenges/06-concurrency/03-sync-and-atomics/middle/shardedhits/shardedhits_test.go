package shardedhits

import (
	"reflect"
	"sync"
	"testing"
)

func TestMeterRecord(t *testing.T) {
	cases := []struct {
		name      string
		shards    int
		recorded  []string
		route     string
		wantCount int64
		wantTotal int64
	}{
		{"single_route", 4, []string{"/orders"}, "/orders", 1, 1},
		{"repeated_route", 4, []string{"/orders", "/orders", "/orders"}, "/orders", 3, 3},
		{"two_routes", 4, []string{"/orders", "/users"}, "/users", 1, 2},
		{"unknown_route", 4, []string{"/orders"}, "/missing", 0, 1},
		{"one_shard", 1, []string{"/a", "/b", "/a"}, "/a", 2, 3},
		{"zero_shards_clamped", 0, []string{"/a", "/a"}, "/a", 2, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := NewMeter(tc.shards)
			for _, r := range tc.recorded {
				m.Record(r)
			}
			if got := m.Count(tc.route); got != tc.wantCount {
				t.Errorf("Count(%q) = %d, want %d", tc.route, got, tc.wantCount)
			}
			if got := m.Total(); got != tc.wantTotal {
				t.Errorf("Total() = %d, want %d", got, tc.wantTotal)
			}
		})
	}
}

func TestMeterRoutesSorted(t *testing.T) {
	m := NewMeter(8)
	for _, r := range []string{"/users", "/orders", "/health", "/orders"} {
		m.Record(r)
	}
	want := []string{"/health", "/orders", "/users"}
	if got := m.Routes(); !reflect.DeepEqual(got, want) {
		t.Errorf("Routes() = %v, want %v", got, want)
	}
}

func TestMeterConcurrentRecord(t *testing.T) {
	m := NewMeter(8)
	routes := []string{"/orders", "/users", "/health", "/search"}
	const workers, perWorker = 8, 200

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < perWorker; j++ {
				m.Record(routes[j%len(routes)])
			}
		}()
	}
	// Readers run against the writers so an unsynchronised map trips -race.
	var readers sync.WaitGroup
	readers.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer readers.Done()
			for j := 0; j < perWorker; j++ {
				_ = m.Count("/orders")
				_ = m.Total()
			}
		}()
	}
	wg.Wait()
	readers.Wait()

	wantPerRoute := int64(workers * perWorker / len(routes))
	for _, r := range routes {
		if got := m.Count(r); got != wantPerRoute {
			t.Errorf("Count(%q) = %d, want %d", r, got, wantPerRoute)
		}
	}
	if got, want := m.Total(), int64(workers*perWorker); got != want {
		t.Errorf("Total() = %d, want %d", got, want)
	}
	if got := len(m.Routes()); got != len(routes) {
		t.Errorf("len(Routes()) = %d, want %d", got, len(routes))
	}
}
