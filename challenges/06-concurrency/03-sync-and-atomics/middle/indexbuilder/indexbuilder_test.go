package indexbuilder

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

// counting returns a build func that records how often it ran per collection.
func counting(calls *sync.Map) func(string) string {
	return func(collection string) string {
		c, _ := calls.LoadOrStore(collection, new(atomic.Int64))
		c.(*atomic.Int64).Add(1)
		return collection + "-idx"
	}
}

func buildCount(calls *sync.Map, collection string) int64 {
	c, ok := calls.Load(collection)
	if !ok {
		return 0
	}
	return c.(*atomic.Int64).Load()
}

func TestRegistryIndex(t *testing.T) {
	cases := []struct {
		name       string
		lookups    []string
		collection string
		want       string
		wantBuilds int64
		wantBuilt  int
	}{
		{"first_use", []string{"orders"}, "orders", "orders-idx", 1, 1},
		{"repeated_use", []string{"orders", "orders", "orders"}, "orders", "orders-idx", 1, 1},
		{"two_collections", []string{"orders", "users"}, "users", "users-idx", 1, 2},
		{"interleaved", []string{"orders", "users", "orders"}, "orders", "orders-idx", 1, 2},
		{"empty_name", []string{""}, "", "-idx", 1, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls sync.Map
			r := NewRegistry(counting(&calls))
			for _, c := range tc.lookups {
				r.Index(c)
			}
			if got := r.Index(tc.collection); got != tc.want {
				t.Errorf("Index(%q) = %q, want %q", tc.collection, got, tc.want)
			}
			if got := buildCount(&calls, tc.collection); got != tc.wantBuilds {
				t.Errorf("build ran %d times for %q, want %d", got, tc.collection, tc.wantBuilds)
			}
			if got := r.Built(); got != tc.wantBuilt {
				t.Errorf("Built() = %d, want %d", got, tc.wantBuilt)
			}
		})
	}
}

func TestRegistryConcurrentSameCollection(t *testing.T) {
	var calls sync.Map
	r := NewRegistry(counting(&calls))

	const workers = 32
	start := make(chan struct{})
	results := make([]string, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			<-start
			results[i] = r.Index("orders")
		}(i)
	}
	close(start)
	wg.Wait()

	if got := buildCount(&calls, "orders"); got != 1 {
		t.Errorf("build ran %d times, want exactly 1", got)
	}
	for i, got := range results {
		if got != "orders-idx" {
			t.Errorf("worker %d saw %q, want %q", i, got, "orders-idx")
		}
	}
}

func TestRegistryConcurrentManyCollections(t *testing.T) {
	var calls sync.Map
	r := NewRegistry(counting(&calls))

	const collections, workersPer = 8, 8
	var wg sync.WaitGroup
	wg.Add(collections * workersPer)
	for c := 0; c < collections; c++ {
		for w := 0; w < workersPer; w++ {
			go func(c int) {
				defer wg.Done()
				name := "c" + strconv.Itoa(c)
				if got := r.Index(name); got != name+"-idx" {
					t.Errorf("Index(%q) = %q", name, got)
				}
			}(c)
		}
	}
	wg.Wait()

	for c := 0; c < collections; c++ {
		name := "c" + strconv.Itoa(c)
		if got := buildCount(&calls, name); got != 1 {
			t.Errorf("build for %q ran %d times, want 1", name, got)
		}
	}
	if got := r.Built(); got != collections {
		t.Errorf("Built() = %d, want %d", got, collections)
	}
}
