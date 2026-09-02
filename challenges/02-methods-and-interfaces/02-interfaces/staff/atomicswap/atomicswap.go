// Package atomicswap — Gopher Workplace challenge.
package atomicswap

import "sync/atomic"

// Policy decides whether a request proceeds.
type Policy interface {
	Allow(key string) bool
}

// AllowAll permits everything.
type AllowAll struct{}

// Allow always returns true.
func (AllowAll) Allow(key string) bool { return true }

// DenyAll rejects everything.
type DenyAll struct{}

// Allow always returns false.
func (DenyAll) Allow(key string) bool { return false }

// PrefixPolicy allows keys with a given prefix.
type PrefixPolicy struct {
	Prefix string
}

// Allow permits keys starting with Prefix.
func (p PrefixPolicy) Allow(key string) bool {
	return len(key) >= len(p.Prefix) && key[:len(p.Prefix)] == p.Prefix
}

// holder wraps a policy so it can live in an atomic.Pointer.
type holder struct {
	p Policy
}

// Strategy holds a hot-swappable policy.
type Strategy struct {
	cur atomic.Pointer[holder]
}

// Set installs a policy.
func (s *Strategy) Set(p Policy) {
	// TODO(candidate): store a pointer to a fully built holder.
	panic("not implemented")
}

// Get returns the current policy, or nil when none is set.
func (s *Strategy) Get() Policy {
	// TODO(candidate): atomic load, nil-safe.
	panic("not implemented")
}

// Allow dispatches to the current policy. With no policy set it fails open.
//
// Examples:
//
//	no policy       => true
//	DenyAll{} set   => false
func (s *Strategy) Allow(key string) bool {
	// TODO(candidate): lock-free dispatch with a fail-open default.
	panic("not implemented")
}
