# Graceful Shutdown Drain

## Intuition

The shutdown goroutine wants to sleep until a condition about shared state becomes true. A channel cannot express "the counter reached zero"; a `Cond` can — it hands back the lock while sleeping and reacquires it on wake, so the predicate is always checked under the lock.

## Approach

1. `NewDrain` creates the struct, then `d.cond = sync.NewCond(&d.mu)`.
2. `Start` locks and increments; `Done` locks, decrements, and `Broadcast`s when the count hits zero.
3. `Wait` locks and loops `for d.inflight > 0 { d.cond.Wait() }`.

## Solution

```go
func NewDrain() *Drain {
	d := &Drain{}
	d.cond = sync.NewCond(&d.mu)
	return d
}

func (d *Drain) Start() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.inflight++
}

func (d *Drain) Done() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.inflight--
	if d.inflight == 0 {
		d.cond.Broadcast()
	}
}

func (d *Drain) Wait() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for d.inflight > 0 {
		d.cond.Wait()
	}
}
```

## Walkthrough

32 handlers are in flight and shutdown calls `Wait`: the predicate holds, so it sleeps and drops the lock. Each `Done` takes the lock and decrements. The one that reaches zero calls `Broadcast`; the waiter wakes, re-acquires the lock, re-tests `inflight > 0`, finds it false, and returns.

## Pitfalls

- `if d.inflight > 0 { d.cond.Wait() }` — a spurious or stale wake-up lets shutdown proceed while requests are still running.
- Calling `Wait` without holding `L` panics.
- Using `Signal` with several waiters leaves the others asleep forever.
