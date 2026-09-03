# Sharded Route Meter

## Intuition

Sharding does not remove the lock, it removes the *queue*. With eight shards, eight goroutines on eight different routes proceed in parallel; only a collision on the same shard costs a wait. The invariant that keeps this correct is that a key's shard is a pure function of the key.

## Approach

1. `Record`: find the shard, lock it, `counts[route]++`, unlock.
2. `Count`: find the shard, lock it, return `counts[route]` (a missing key yields the zero value).
3. `Total`: walk the shards, locking and unlocking one at a time, summing the values.
4. `Routes`: collect keys the same way, then `sort.Strings`.

## Solution

```go
// Record adds one hit for a route, locking only that route's shard.
//
// Examples:
//
//	m := NewMeter(4); m.Record("/orders"); m.Count("/orders") => 1
func (m *Meter) Record(route string) {
	s := m.shardFor(route)
	s.mu.Lock()
	s.counts[route]++
	s.mu.Unlock()
}

// Count returns the hits recorded for a route, or 0 if it was never seen.
//
// Examples:
//
//	NewMeter(4).Count("/unknown") => 0
func (m *Meter) Count(route string) int64 {
	s := m.shardFor(route)
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.counts[route]
}

// Total returns the hits across every route and shard.
//
// Examples:
//
//	m.Record("/a"); m.Record("/b"); m.Total() => 2
func (m *Meter) Total() int64 {
	var total int64
	for _, s := range m.shards {
		s.mu.Lock()
		for _, c := range s.counts {
			total += c
		}
		s.mu.Unlock()
	}
	return total
}

// Routes returns every recorded route, sorted.
//
// Examples:
//
//	m.Record("/b"); m.Record("/a"); m.Routes() => ["/a" "/b"]
func (m *Meter) Routes() []string {
	var routes []string
	for _, s := range m.shards {
		s.mu.Lock()
		for r := range s.counts {
			routes = append(routes, r)
		}
		s.mu.Unlock()
	}
	sort.Strings(routes)
	return routes
}
```

## Walkthrough

- `Record("/orders")` hashes to one shard; a concurrent `Record("/users")` on a different shard never blocks.
- `Count` on an unrecorded route returns the map's zero value, `0`, without inserting anything.
- `Total` releases each shard lock before taking the next, so it can never deadlock against a writer.
- In the concurrency test 8 writers and 2 readers run together; without the per-shard lock, `-race` reports a concurrent map write.

## Pitfalls

- Reading `s.counts[route]` without the lock — a map read racing a map write is a data race, and Go may crash outright with "concurrent map read and map write".
- Holding every shard lock at once in `Total`; one at a time is enough and avoids lock-order problems.
- Choosing the shard from anything but the key (round-robin, a counter) — the same route then lands in two shards and the count is wrong.
