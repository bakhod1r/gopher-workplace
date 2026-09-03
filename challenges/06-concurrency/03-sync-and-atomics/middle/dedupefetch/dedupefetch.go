// Package dedupefetch — Gopher Workplace challenge.
package dedupefetch

import "sync"

// Fetcher caches catalogue lookups so a key is loaded from upstream exactly
// once, no matter how many goroutines ask for it.
type Fetcher struct {
	mu      sync.Mutex
	entries map[string]*entry
}

type entry struct {
	once  sync.Once
	value string
}

// NewFetcher returns an empty Fetcher.
//
// Examples:
//
//	NewFetcher().Len() => 0
func NewFetcher() *Fetcher {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Fetch returns the value for key, calling load exactly once per key even
// when many goroutines request the same key at the same time.
//
// Examples:
//
//	f.Fetch("sku-1", load) => "product:sku-1"
//	f.Fetch("sku-1", load) => "product:sku-1" (load not called again)
//	f.Fetch("sku-2", load) => "product:sku-2"
func (f *Fetcher) Fetch(key string, load func(key string) string) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Len returns how many keys have been requested.
func (f *Fetcher) Len() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.entries)
}
