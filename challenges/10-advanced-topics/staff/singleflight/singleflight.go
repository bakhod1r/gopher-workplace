// Package singleflight — Gopher Workplace challenge.
package singleflight

import "sync"

// call is one in-flight or completed fetch.
type call struct {
	wg  sync.WaitGroup
	val int
}

// Group deduplicates concurrent calls by key.
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do runs fn for key and returns its result, sharing one in-flight call
// among every concurrent caller for that key.
//
// A thundering herd on a cold cache must not become N identical fetches,
// each allocating its own result.
//
// Examples:
//
//	g.Do("a", expensive) from 32 goroutines => expensive runs once
func (g *Group) Do(key string, fn func() int) int {
	panic("not implemented")
}
