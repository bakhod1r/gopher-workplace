# Read-Copy-Update

## Intuition

RCU inverts the usual cost model. A `RWMutex` makes readers cheap but not free —
they still write to the lock's internal counter, bouncing a cache line between
cores. RCU makes reads a single atomic load of an immutable object: no shared
writes at all, so read throughput scales linearly with cores.

The price is that writers must copy, and that old versions stay alive until the
last reader releases them.

## Approach

1. Serialize writers with a plain mutex.
2. Load the current version.
3. Construct a new object — never touch the old one.
4. Publish it atomically.

## Solution

```go
func (r *RCU) Update(newData string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	old := r.ptr.Load()
	next := &Config{Data: newData}
	_ = old // a real update derives `next` from `old`
	r.ptr.Store(next)
}
```

## Walkthrough

`New` publishes `&Config{Data: "v1"}`. `Update("v2")` allocates a second
`Config` and stores its address. A reader that loaded the pointer before the
store still holds the `"v1"` object, which is untouched and stays valid for as
long as that reader keeps the pointer — Go's garbage collector is what frees it,
which is why Go needs no explicit grace period.

The mutex never touches the read path. Two concurrent writers, though, would
otherwise both derive from the same `old` and one update would vanish — which is
what it prevents.

## Pitfalls

- **`cfg := r.ptr.Load(); cfg.Data = newData`.** The obvious shortcut and a
  straight data race: readers are dereferencing that object right now.
- **Dropping the mutex.** With a single field it looks harmless; as soon as an
  update is read-modify-write, two writers lose an update.
- **Expecting readers to see the update immediately.** A reader that already
  loaded the pointer keeps the old version, by design.
- **A `RWMutex` in the read path.** That is the design RCU exists to replace.

## Why Go makes RCU easy

In C, the hard part of RCU is knowing when the last reader is done so the old
object can be freed — hence grace periods and `synchronize_rcu`. Go's GC answers
that for free: the old `Config` is collected once no goroutine holds a pointer
to it.
