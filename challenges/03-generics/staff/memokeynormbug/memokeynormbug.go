// Package memokeynormbug — Gopher Workplace challenge.
package memokeynormbug

import (
	"strings"
)

// Key identifies an upstream.
type Key struct {
	Host string
	Port int
}

// Memo caches one value per normalised Key.
type Memo[V any] struct {
	m map[Key]V
}

// Get returns the cached value for k, computing it with fn on a miss.
// Keys are normalised, so equivalent keys share one entry.
//
// Examples:
//
//	Get(Key{"API.example.com ", 443}, fn) // same entry as Key{"api.example.com", 443}
func (m *Memo[V]) Get(k Key, fn func(Key) V) V {
	// CHANGE CODE BELOW THIS LINE
	if m.m == nil {
		m.m = make(map[Key]V)
	}
	if v, ok := m.m[k]; ok {
		return v
	}
	v := fn(k)
	m.m[k] = v
	return v
	// CHANGE CODE ABOVE THIS LINE
}

// Norm canonicalises a key: the host is trimmed and lower-cased.
func Norm(k Key) Key {
	k.Host = strings.ToLower(strings.TrimSpace(k.Host))
	return k
}
