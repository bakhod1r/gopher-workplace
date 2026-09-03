// Package routetable — Gopher Workplace challenge.
package routetable

import "sync"

// Table maps request paths to backend addresses. Every request reads it; the
// config watcher writes to it only when a deploy changes a route.
type Table struct {
	mu     sync.RWMutex
	routes map[string]string
}

// NewTable returns an empty routing table.
//
// Examples:
//
//	NewTable().Len() => 0
func NewTable() *Table {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Set points a path at a backend, replacing any previous target.
//
// Examples:
//
//	t := NewTable(); t.Set("/api", "backend-1"); t.Len() => 1
func (t *Table) Set(path, backend string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Lookup returns the backend for a path.
//
// Examples:
//
//	t.Set("/api", "backend-1"); t.Lookup("/api")  => "backend-1", true
//	NewTable().Lookup("/missing")                 => "", false
func (t *Table) Lookup(path string) (string, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Snapshot returns a copy of the table that the caller may keep and read
// without any further locking.
//
// Examples:
//
//	t.Set("/api", "backend-1"); t.Snapshot() => map[/api:backend-1]
func (t *Table) Snapshot() map[string]string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Len returns the number of routes.
func (t *Table) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.routes)
}
