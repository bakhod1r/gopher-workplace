package connregistry

import (
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestRegistry(t *testing.T) {
	cases := []struct {
		name     string
		register [][2]string
		lookup   string
		wantAddr string
		wantOK   bool
		wantIDs  string
	}{
		{"registered", [][2]string{{"a", "10.0.0.1"}}, "a", "10.0.0.1", true, "a"},
		{"unknown", [][2]string{{"a", "10.0.0.1"}}, "ghost", "", false, "a"},
		{"empty", nil, "a", "", false, ""},
		{"sorted_ids", [][2]string{{"b", "x"}, {"a", "y"}}, "b", "x", true, "a,b"},
		{"readdress", [][2]string{{"a", "x"}, {"a", "y"}}, "a", "y", true, "a"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRegistry()
			for _, e := range tc.register {
				r.Register(e[0], e[1])
			}
			addr, ok := r.Lookup(tc.lookup)
			if addr != tc.wantAddr || ok != tc.wantOK {
				t.Errorf("Lookup(%q) = %q, %v, want %q, %v", tc.lookup, addr, ok, tc.wantAddr, tc.wantOK)
			}
			ids := r.IDs()
			sort.Strings(ids)
			if got := strings.Join(ids, ","); got != tc.wantIDs {
				t.Errorf("IDs() = %q, want %q", got, tc.wantIDs)
			}
		})
	}
}

func TestRegistryConcurrent(t *testing.T) {
	r := NewRegistry()
	for i := 0; i < 10; i++ {
		r.Register(strconv.Itoa(i), "addr")
	}

	var wg sync.WaitGroup
	wg.Add(24)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				r.Lookup(strconv.Itoa(j % 10))
				r.IDs()
			}
		}()
	}
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				r.Register(strconv.Itoa(j%10), "addr")
			}
		}()
	}
	wg.Wait()

	if got := len(r.IDs()); got != 10 {
		t.Errorf("len(IDs()) = %d, want 10", got)
	}
}
