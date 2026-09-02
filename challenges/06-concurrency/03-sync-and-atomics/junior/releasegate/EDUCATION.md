# Deploy Release Gate

## Intuition

A condition variable is a waiting room attached to a mutex. `Wait` puts you in the room and hands back the lock so the state can change; `Broadcast` empties the room. Because the state is a sticky flag, an already-open gate lets `Wait` fall straight through the loop.

## Approach

1. `Wait` takes the lock and loops on `!g.open`, calling `Wait` inside.
2. `Open` takes the lock, sets `open = true`, releases it, then calls `Broadcast`.
3. `IsOpen` reads the flag under the lock.

## Solution

```go
// Package releasegate - Gopher Workplace challenge.
package releasegate

import "sync"

// Gate holds goroutines until a release opens it.
type Gate struct {
	mu   sync.Mutex
	cond *sync.Cond
	open bool
}

// NewGate returns a closed gate.
func NewGate() *Gate {
	g := &Gate{}
	g.cond = sync.NewCond(&g.mu)
	return g
}

// Wait blocks until the gate is open.
//
// Examples:
//
//	g := NewGate(); g.Open(); g.Wait() => returns immediately
func (g *Gate) Wait() {
	g.mu.Lock()
	defer g.mu.Unlock()
	for !g.open {
		g.cond.Wait()
	}
}

// Open opens the gate and releases every waiting goroutine.
//
// Examples:
//
//	g := NewGate(); g.Open(); g.IsOpen() => true
func (g *Gate) Open() {
	g.mu.Lock()
	g.open = true
	g.mu.Unlock()
	g.cond.Broadcast()
}

// IsOpen reports whether the gate is open.
//
// Examples:
//
//	g := NewGate(); g.IsOpen() => false
func (g *Gate) IsOpen() bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.open
}
```

## Walkthrough

Eight handlers call `Wait` and park. The release calls `Open`: the flag flips and `Broadcast` wakes all eight. Each re-acquires the lock in turn, re-checks `!g.open` (now false), and proceeds.

## Pitfalls

- Using `Signal`, which releases only one handler and hangs the rest.
- Using `if` instead of `for` around `Wait`.
- Reading `g.open` in `IsOpen` without the lock — a race with `Open`.
