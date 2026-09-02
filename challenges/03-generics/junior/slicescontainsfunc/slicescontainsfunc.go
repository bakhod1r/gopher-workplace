// Package slicescontainsfunc — Gopher Workplace challenge.
package slicescontainsfunc

// Entry is a cache entry with a time-to-live in seconds.
type Entry struct {
	Key string
	TTL int
}

// AnyExpired reports whether any entry has ttl <= 0.
func AnyExpired(entries []Entry) bool {
	// TODO(candidate): use the stdlib predicate search.
	panic("not implemented")
}
