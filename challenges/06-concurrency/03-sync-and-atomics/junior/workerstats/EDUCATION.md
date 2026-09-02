# Worker Pool Snapshot

## Intuition

Correctness here is not about any single field but about the pair. Holding one read lock across both loads freezes the state for the duration of the read, and returning a copy carries that frozen instant out of the critical section.

## Approach

1. `Start` and `Fail` write-lock and increment their counter.
2. `Snapshot` read-locks, builds the struct from both fields, and returns it.
3. Use `defer` so the unlock happens after the struct is built.

## Solution

```go
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
	p.mu.Lock()
	p.started++
	p.mu.Unlock()
}

// Fail records that a job failed.
//
// Examples:
//
//	var p Pool; p.Start(); p.Fail(); p.Snapshot() => Snapshot{Started: 1, Failed: 1}
func (p *Pool) Fail() {
	p.mu.Lock()
	p.failed++
	p.mu.Unlock()
}

// Snapshot returns both counters read at a single instant.
//
// Examples:
//
//	var p Pool; p.Snapshot() => Snapshot{0, 0}
func (p *Pool) Snapshot() Snapshot {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return Snapshot{Started: p.started, Failed: p.failed}
}
```

## Walkthrough

The stats endpoint takes `RLock` and reads `started = 5, failed = 2`. A worker's `Fail` waits for the write lock, so the endpoint's snapshot reflects one exact moment rather than a blend of two.

## Pitfalls

- Implementing `Snapshot` as `Snapshot{p.Started(), p.Failed()}` — two lock holds, two different instants.
- Returning a `*Snapshot` pointing at shared state.
- Using `RLock` in `Start`, which does not exclude another writer.
