// Package warmupgate — Gopher Workplace challenge.
package warmupgate

// WarmCaches warms every key concurrently and returns a single error for the
// deploy gate: the failure belonging to the lowest-indexed key that could not be
// warmed, or nil when the whole set is hot. Every warm is allowed to finish
// first, so the pod never starts serving with half-warm caches and never leaks
// a goroutine into the next release.
//
// Examples:
//
//	WarmCaches([]string{"tags", "prices"}, warm)  => <nil>
//	WarmCaches([]string{"tags", "bad"}, warm)     => errColdSource
//	WarmCaches(nil, warm)                         => <nil>
func WarmCaches(keys []string, warm func(key string) error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
