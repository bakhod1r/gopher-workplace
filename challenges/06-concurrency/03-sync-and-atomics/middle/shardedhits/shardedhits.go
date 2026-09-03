// Package shardedhits — Gopher Workplace challenge.
package shardedhits

import (
	"hash/fnv"
	"sync"
)

// Meter counts requests per route at the API gateway. One global mutex would
// serialise every worker on every request, so the counts are split across a
// fixed number of shards and each shard carries its own lock.
type Meter struct {
	shards []*routeShard
}

type routeShard struct {
	mu     sync.Mutex
	counts map[string]int64
}

// NewMeter returns a Meter with n shards; n <= 0 collapses to a single shard.
//
// Examples:
//
//	NewMeter(8).Total() => 0
func NewMeter(n int) *Meter {
	if n <= 0 {
		n = 1
	}
	m := &Meter{shards: make([]*routeShard, n)}
	for i := range m.shards {
		m.shards[i] = &routeShard{counts: make(map[string]int64)}
	}
	return m
}

// shardFor returns the shard that owns a route.
func (m *Meter) shardFor(route string) *routeShard {
	h := fnv.New32a()
	_, _ = h.Write([]byte(route))
	return m.shards[int(h.Sum32())%len(m.shards)]
}

// Record adds one hit for a route, locking only that route's shard.
//
// Examples:
//
//	m := NewMeter(4); m.Record("/orders"); m.Count("/orders") => 1
func (m *Meter) Record(route string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Count returns the hits recorded for a route, or 0 if it was never seen.
//
// Examples:
//
//	NewMeter(4).Count("/unknown") => 0
func (m *Meter) Count(route string) int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Total returns the hits across every route and shard.
//
// Examples:
//
//	m.Record("/a"); m.Record("/b"); m.Total() => 2
func (m *Meter) Total() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Routes returns every recorded route, sorted.
//
// Examples:
//
//	m.Record("/b"); m.Record("/a"); m.Routes() => ["/a" "/b"]
func (m *Meter) Routes() []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
