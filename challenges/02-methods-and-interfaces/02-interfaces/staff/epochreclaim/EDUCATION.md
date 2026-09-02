# Epoch Reclamation

## Intuition

Reclamation is safe when no goroutine can still hold a reference. Epochs give a cheap conservative answer: a reader that entered *after* the retirement never saw the object, so only readers from earlier epochs matter.

## Approach

1. `Enter` records a reader in the current epoch and returns it.
2. `Exit` decrements, removing the entry at zero so the map stays small.
3. `Retire` records the object with the current epoch, then advances — later readers are automatically excluded.
4. `Reclaim` frees objects with no active reader at or before their epoch, compacting `pending` in place.

## Solution

```go
func (e *Epoch) Enter() uint64 {
	e.mu.Lock()
	defer e.mu.Unlock()

	ep := e.current
	e.active[ep]++
	return ep
}

func (e *Epoch) Retire(obj Freer) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.pending = append(e.pending, retired{obj: obj, epoch: e.current})
	e.current++
}

func (e *Epoch) Reclaim() int {
	e.mu.Lock()
	defer e.mu.Unlock()

	freed := 0
	kept := e.pending[:0]
	for _, r := range e.pending {
		if e.hasReaderAtOrBefore(r.epoch) {
			kept = append(kept, r)
			continue
		}
		r.obj.Free()
		freed++
	}
	e.pending = kept
	return freed
}
```

## Walkthrough

`TestLaterReaderDoesNotBlockReclaim` is the point of advancing the epoch on retire: the new reader entered epoch 1, the node was retired in epoch 0, so it can be freed even with a reader active.

## Pitfalls

- Freeing immediately on retire — the classic use-after-free in lock-free code.
- Not advancing the epoch, so every later reader blocks reclamation forever.
- Leaving zero-count entries in `active`, which makes `hasReaderAtOrBefore` report phantom readers.
