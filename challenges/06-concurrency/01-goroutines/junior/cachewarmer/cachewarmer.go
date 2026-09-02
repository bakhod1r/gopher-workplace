// Package cachewarmer — Gopher Workplace challenge.
package cachewarmer

// WarmAll loads every key and reports the cached size, clamped to the entry cap.
//
// Examples:
//
//	WarmAll([]string{"a", "bb"}, loader, 1000)  => [100 200]
//	WarmAll([]string{"huge"}, loader, 150)      => [150]
//	WarmAll(nil, loader, 100)                   => []
func WarmAll(keys []string, load func(key string) int, capBytes int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
