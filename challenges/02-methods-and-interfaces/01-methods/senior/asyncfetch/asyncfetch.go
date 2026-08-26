// Package asyncfetch — Gopher Workplace challenge.
package asyncfetch

import "sync"

// Fetcher retrieves data.
type Fetcher struct {
	Result string
	mu     sync.Mutex
}

// Fetch sets Result to "data: " + id.
func (f *Fetcher) Fetch(id string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Result = "data: " + id
}

// FetchAsync runs Fetch in a new goroutine.
//
// Examples:
//
//	f := &Fetcher{}
//	done := f.FetchAsync("123")
//	<-done // waits for completion
func (f *Fetcher) FetchAsync(id string) <-chan struct{} {
	// TODO(candidate): launch a goroutine that calls Fetch(id), then closes
	// the returned channel to signal completion.
	panic("not implemented")
}
