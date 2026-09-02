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
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Allowed reports whether ip is permitted.
//
// Examples:
//
//	NewAllowlist().Allowed("10.0.0.9") => false
func (a *Allowlist) Allowed(ip string) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Size returns the number of allowed IPs.
//
// Examples:
//
//	a.Allow("10.0.0.1"); a.Allow("10.0.0.2"); a.Size() => 2
func (a *Allowlist) Size() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
