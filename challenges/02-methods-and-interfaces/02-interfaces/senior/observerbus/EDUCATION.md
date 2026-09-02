# Observer Bus

## Intuition

Delivering while holding the lock invites deadlock the moment a handler touches the bus. Snapshotting first decouples the two: mutations during delivery affect the *next* publish, not the one in flight.

## Approach

1. `Subscribe` allocates an id under the lock and returns a closure that deletes it.
2. `delete` on a missing key is a no-op, which makes unsubscribe idempotent for free.
3. `Publish` copies the handlers into a slice under the lock, unlocks, and then calls them.
4. `Count` reads the map size under the lock.

## Solution

```go
func (b *Bus) Subscribe(h Handler) func() {
	b.mu.Lock()
	id := b.nextID
	b.nextID++
	b.subs[id] = h
	b.mu.Unlock()

	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		delete(b.subs, id)
	}
}

func (b *Bus) Publish(event string) {
	b.mu.Lock()
	snapshot := make([]Handler, 0, len(b.subs))
	for _, h := range b.subs {
		snapshot = append(snapshot, h)
	}
	b.mu.Unlock()

	for _, h := range snapshot {
		h.Handle(event)
	}
}
```

## Walkthrough

`TestUnsubscribeDuringPublish` has a handler remove itself mid-delivery. Because delivery walks a snapshot and the lock is free, the `delete` succeeds and the second publish skips it.

## Pitfalls

- Ranging over `b.subs` while holding the lock and calling handlers — a self-unsubscribing handler deadlocks on the same mutex.
- Iterating the map without a snapshot, so concurrent mutation is a race.
- Using slice indices as subscription ids, which shift when other subscribers are removed.
