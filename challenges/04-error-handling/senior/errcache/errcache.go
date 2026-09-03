// Package errcache — Gopher Workplace challenge.
package errcache

// entry holds a completed load result.
type entry struct {
	v   int
	err error
}

// Cache memoizes results of Load, including failures.
type Cache struct {
	Load func(string) (int, error)

	entries map[string]entry
}

// Get returns the cached result for key, loading it at most once.
//
// Examples:
//
//	c.Get("a") => 1, nil
func (c *Cache) Get(key string) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
