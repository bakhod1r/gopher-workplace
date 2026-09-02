# Firewall Allowlist

## Intuition

`Allow` is a check-then-act: two admins adding the same IP must not both be told they added it. Doing the lookup and the insert in one lock hold makes the pair indivisible.

## Approach

1. `NewAllowlist` makes the map.
2. `Allow` locks, checks membership, inserts when absent, returns whether it inserted.
3. `Allowed` and `Size` read under the lock.

## Solution

```go
// Package allowlist - Gopher Workplace challenge.
package allowlist

import "sync"

// Allowlist is a concurrency-safe set of permitted client IPs.
type Allowlist struct {
	mu  sync.Mutex
	ips map[string]struct{}
}

// NewAllowlist returns an empty allowlist.
func NewAllowlist() *Allowlist {
	return &Allowlist{ips: make(map[string]struct{})}
}

// Allow adds ip and reports whether it was newly added.
//
// Examples:
//
//	a := NewAllowlist(); a.Allow("10.0.0.1") => true
//	a.Allow("10.0.0.1")                      => false
func (a *Allowlist) Allow(ip string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.ips[ip]; ok {
		return false
	}
	a.ips[ip] = struct{}{}
	return true
}

// Allowed reports whether ip is permitted.
//
// Examples:
//
//	NewAllowlist().Allowed("10.0.0.9") => false
func (a *Allowlist) Allowed(ip string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	_, ok := a.ips[ip]
	return ok
}

// Size returns the number of allowed IPs.
//
// Examples:
//
//	a.Allow("10.0.0.1"); a.Allow("10.0.0.2"); a.Size() => 2
func (a *Allowlist) Size() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return len(a.ips)
}
```

## Walkthrough

Two admins call `Allow("10.0.0.1")` at once. The first takes the lock, finds nothing, inserts, returns true. The second then finds the entry and returns false — one insert, one honest answer.

## Pitfalls

- Checking membership, unlocking, then inserting — both callers can report a new insert.
- Using `map[string]bool` and forgetting that a stored `false` still counts as present.
- Forgetting to lock `Allowed`, which races with the admin's insert.
