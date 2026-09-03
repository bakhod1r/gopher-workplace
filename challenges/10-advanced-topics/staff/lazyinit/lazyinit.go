// Package lazyinit — Gopher Workplace challenge.
package lazyinit

import (
	"sync"
	"sync/atomic"
)

// Builds counts how many times the index has been constructed.
var Builds atomic.Int64

// Table indexes a slice of pairs lazily.
type Table struct {
	once  sync.Once
	pairs [][2]string
	index map[string]int
}

// NewTable returns a table over pairs, without building the index.
func NewTable(pairs [][2]string) *Table {
	return &Table{pairs: pairs}
}

// build constructs the index. It must run at most once per table.
func (t *Table) build() {
	Builds.Add(1)
	t.index = make(map[string]int, len(t.pairs))
	for i, p := range t.pairs {
		t.index[p[0]] = i
	}
}

// Lookup returns the value for k, building the table's index on first use.
//
// The index is expensive, the callers are concurrent, and it must be built
// exactly once — every later lookup should be a plain map read.
//
// Examples:
//
//	t := NewTable(pairs); t.Lookup("a") => the value for "a"
func (t *Table) Lookup(k string) (int, bool) {
	panic("not implemented")
}
