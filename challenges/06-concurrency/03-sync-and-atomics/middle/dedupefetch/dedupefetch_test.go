package dedupefetch

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFetch(t *testing.T) {
	load := func(k string) string { return "product:" + k }

	cases := []struct {
		name string
		keys []string
		key  string
		want string
	}{
		{"single", []string{"sku-1"}, "sku-1", "product:sku-1"},
		{"repeat", []string{"sku-1", "sku-1"}, "sku-1", "product:sku-1"},
		{"two_keys", []string{"sku-1", "sku-2"}, "sku-2", "product:sku-2"},
		{"empty_key", []string{""}, "", "product:"},
		{"many", []string{"a", "b", "c"}, "b", "product:b"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := NewFetcher()
			var got string
			for _, k := range tc.keys {
				v := f.Fetch(k, load)
				if k == tc.key {
					got = v
				}
			}
			if got != tc.want {
				t.Errorf("Fetch(%q) = %q, want %q", tc.key, got, tc.want)
			}
		})
	}
}

func TestLoadRunsOncePerKey(t *testing.T) {
	f := NewFetcher()
	var calls atomic.Int64
	load := func(k string) string {
		calls.Add(1)
		return "product:" + k
	}

	for range 5 {
		f.Fetch("sku-1", load)
	}
	f.Fetch("sku-2", load)

	if got := calls.Load(); got != 2 {
		t.Errorf("load called %d times, want 2", got)
	}
	if got := f.Len(); got != 2 {
		t.Errorf("Len() = %d, want 2", got)
	}
}

func TestConcurrentFetchesCollapse(t *testing.T) {
	f := NewFetcher()
	var calls atomic.Int64
	load := func(k string) string {
		calls.Add(1)
		return "product:" + k
	}

	const keys, callers = 4, 200
	var wg sync.WaitGroup
	for i := range callers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "sku-" + strconv.Itoa(i%keys)
			if got, want := f.Fetch(key, load), "product:"+key; got != want {
				t.Errorf("Fetch(%q) = %q, want %q", key, got, want)
			}
		}(i)
	}
	wg.Wait()

	if got := calls.Load(); got != keys {
		t.Errorf("load called %d times, want exactly %d", got, keys)
	}
}
