// Package workerstats - Gopher Workplace challenge.
package workerstats

import "sync"

// Snapshot is a consistent view of the pool counters.
type Snapshot struct {
	Started int
	Failed  int
}

// Pool records worker events and serves consistent stats snapshots.
type Pool struct {
	mu      sync.RWMutex
	started int
	failed  int
}

// Start records that a job started.
//
// Examples:
//
//	var p Pool; p.Start(); p.Snapshot() => Snapshot{Started: 1, Failed: 0}
func (p *Pool) Start() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Fail records that a job failed.
//
// Examples:
//
//	var p Pool; p.Start(); p.Fail(); p.Snapshot() => Snapshot{Started: 1, Failed: 1}
func (p *Pool) Fail() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Snapshot returns both counters read at a single instant.
//
// Examples:
//
//	var p Pool; p.Snapshot() => Snapshot{0, 0}
func (p *Pool) Snapshot() Snapshot {
	// TODO(candidate): implement this.
	panic("not implemented")
}
