# Object Store Upload Window

## Intuition

A permit window is a counter plus a waiting room. The counter needs a lock; the waiting room needs a way to sleep without holding that lock. `sync.Cond` is exactly the pair: `Wait` puts you to sleep *and* hands the lock back.

## Approach

1. Clamp `limit` to at least 1 and build the `Cond` over the struct's mutex.
2. `Acquire`: lock, `for inUse >= limit { cond.Wait() }`, then `inUse++`.
3. `Release`: lock, `inUse--`, `cond.Broadcast()`.

## Solution

```go
func NewWindow(limit int) *Window {
	if limit < 1 {
		limit = 1
	}
	w := &Window{limit: limit}
	w.cond = sync.NewCond(&w.mu)
	return w
}

func (w *Window) Acquire() {
	w.mu.Lock()
	defer w.mu.Unlock()
	for w.inUse >= w.limit {
		w.cond.Wait()
	}
	w.inUse++
}

func (w *Window) Release() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.inUse--
	w.cond.Broadcast()
}
```

## Walkthrough

With limit 4 and 300 uploaders, the first four increment and run. The fifth finds `inUse == 4`, sleeps and drops the lock. When any uploader releases, the broadcast wakes every sleeper; each re-acquires the lock, re-tests the predicate, and exactly as many as there are free permits get through — the rest go back to sleep.

## Pitfalls

- `if` instead of `for`: a woken goroutine may find the window full again and takes a permit anyway.
- Broadcasting before decrementing — waiters wake, see a full window, and sleep again for no reason.
- Calling `Wait` without holding the lock: that panics.
