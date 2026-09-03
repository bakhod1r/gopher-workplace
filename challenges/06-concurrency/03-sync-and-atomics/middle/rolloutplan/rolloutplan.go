// Package rolloutplan — Gopher Workplace challenge.
package rolloutplan

import (
	"sync"
	"sync/atomic"
)

// Plan holds the rollout percentage of every feature flag. Request handlers
// read it on every request; the config watcher rewrites it a few times a day.
// Writers therefore build a whole new map and publish it atomically, so
// readers never take a lock at all.
type Plan struct {
	writers sync.Mutex // serialises writers so two updates cannot lose each other
	current atomic.Pointer[map[string]int]
}

// NewPlan returns a Plan seeded with a copy of initial. Percentages are
// clamped into 0..100.
//
// Examples:
//
//	NewPlan(map[string]int{"checkout": 25}).Percent("checkout") => 25
//	NewPlan(nil).Percent("checkout")                            => 0
func NewPlan(initial map[string]int) *Plan {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// clamp confines a percentage to 0..100.
func clamp(pct int) int {
	if pct < 0 {
		return 0
	}
	if pct > 100 {
		return 100
	}
	return pct
}

// Percent returns a feature's rollout percentage, or 0 if it is unknown.
// It takes no lock: it reads whichever snapshot is current.
//
// Examples:
//
//	NewPlan(nil).Percent("unknown") => 0
func (p *Plan) Percent(feature string) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// SetPercent publishes a new snapshot with feature set to pct, clamped into
// 0..100. The previous snapshot is left untouched for readers still using it.
//
// Examples:
//
//	p := NewPlan(nil); p.SetPercent("checkout", 40); p.Percent("checkout") => 40
//	p.SetPercent("checkout", 300); p.Percent("checkout")                   => 100
func (p *Plan) SetPercent(feature string, pct int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Features returns every feature in the current snapshot, sorted.
//
// Examples:
//
//	p.SetPercent("b", 1); p.SetPercent("a", 1); p.Features() => ["a" "b"]
func (p *Plan) Features() []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
