# Finalizer Queue

## Intuition

A finalizer is a safety net, not a destructor: it runs at an unpredictable time, only if the object becomes unreachable, and only once. Correct code closes explicitly and clears the finalizer so the net never fires.

## Approach

1. `NewHandle` builds the handle and registers a finalizer calling the shared `release`.
2. `release` uses `sync.Once` so explicit and finalizer paths cannot double-release.
3. `Close` releases, then calls `runtime.SetFinalizer(h, nil)` to unregister.
4. `Closed` reads an atomic flag, safe from the finalizer goroutine.

## Solution

```go
func NewHandle(p Resource) *Handle {
	h := &Handle{pool: p}
	runtime.SetFinalizer(h, func(h *Handle) {
		h.release()
	})
	return h
}

func (h *Handle) release() {
	h.once.Do(func() {
		h.closed.Store(true)
		h.pool.Release()
	})
}

func (h *Handle) Close() {
	h.release()
	runtime.SetFinalizer(h, nil)
}
```

## Walkthrough

Finalizers run on a runtime goroutine, so `release` must be safe concurrently with the owner's `Close` — hence the `sync.Once` and the atomic flag rather than a plain bool.

## Pitfalls

- Capturing the handle itself in the finalizer closure, which makes it permanently reachable and the finalizer never runs.
- Relying on finalizers for ordinary cleanup: they may not run before exit, and their timing is not a contract.
- Forgetting `SetFinalizer(h, nil)` in `Close` — harmless here thanks to `Once`, but it keeps the object in the finalizer queue for an extra GC cycle.
