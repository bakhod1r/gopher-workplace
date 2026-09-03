package profilecache

import (
	"strconv"
	"sync"
	"testing"
)

func TestCacheGet(t *testing.T) {
	cases := []struct {
		name       string
		seed       map[string]string
		lookups    []string
		want       string
		wantOK     bool
		wantHits   int64
		wantMisses int64
	}{
		{"hit", map[string]string{"u1": "ada"}, []string{"u1"}, "ada", true, 1, 0},
		{"miss", nil, []string{"u1"}, "", false, 0, 1},
		{"repeated_hits", map[string]string{"u1": "ada"}, []string{"u1", "u1", "u1"}, "ada", true, 3, 0},
		{"mixed", map[string]string{"u1": "ada"}, []string{"u1", "u2", "u1"}, "ada", true, 2, 1},
		{"empty_profile_still_hit", map[string]string{"u1": ""}, []string{"u1"}, "", true, 1, 0},
		{"overwritten", map[string]string{"u1": "grace"}, []string{"u1"}, "grace", true, 1, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCache()
			for k, v := range tc.seed {
				c.Put(k, v)
			}
			var got string
			var ok bool
			for _, id := range tc.lookups {
				got, ok = c.Get(id)
			}
			if got != tc.want || ok != tc.wantOK {
				t.Errorf("Get() = %q, %v; want %q, %v", got, ok, tc.want, tc.wantOK)
			}
			hits, misses := c.Stats()
			if hits != tc.wantHits || misses != tc.wantMisses {
				t.Errorf("Stats() = %d, %d; want %d, %d", hits, misses, tc.wantHits, tc.wantMisses)
			}
		})
	}
}

func TestCacheInvalidate(t *testing.T) {
	c := NewCache()
	c.Put("u1", "ada")
	if !c.Invalidate("u1") {
		t.Error("Invalidate(present) = false, want true")
	}
	if c.Invalidate("u1") {
		t.Error("Invalidate(absent) = true, want false")
	}
	if _, ok := c.Get("u1"); ok {
		t.Error("Get after Invalidate returned a value")
	}
	if got := c.Len(); got != 0 {
		t.Errorf("Len() = %d, want 0", got)
	}
}

func TestCacheConcurrentReadMostly(t *testing.T) {
	c := NewCache()
	const keys = 16
	for i := 0; i < keys; i++ {
		c.Put("u"+strconv.Itoa(i), "profile-"+strconv.Itoa(i))
	}

	const readers, writers, perGoroutine = 8, 2, 300
	var wg sync.WaitGroup
	wg.Add(readers + writers)
	for r := 0; r < readers; r++ {
		go func() {
			defer wg.Done()
			for i := 0; i < perGoroutine; i++ {
				c.Get("u" + strconv.Itoa(i%keys))
			}
		}()
	}
	for w := 0; w < writers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < perGoroutine; i++ {
				c.Put("u"+strconv.Itoa(i%keys), "refreshed")
				c.Invalidate("temp" + strconv.Itoa(w))
			}
		}(w)
	}
	wg.Wait()

	hits, misses := c.Stats()
	if hits+misses != int64(readers*perGoroutine) {
		t.Errorf("hits+misses = %d, want %d", hits+misses, readers*perGoroutine)
	}
	if misses != 0 {
		t.Errorf("misses = %d, want 0 (every key was seeded)", misses)
	}
	if got := c.Len(); got != keys {
		t.Errorf("Len() = %d, want %d", got, keys)
	}
}
