# GC Sweep Phase

## Intuition

Sweeping is the cheap, dumb half of mark-and-sweep: no graph, no recursion, just
a straight walk over the heap freeing everything the mark phase did not touch.
Its cost scales with the size of the heap, not with the size of the live set —
which is exactly why large mostly-garbage heaps are expensive to collect.

## Approach

1. Walk every slot.
2. Count the unmarked ones.

## Solution

```go
func (h *Heap) Sweep() int {
	freed := 0
	for _, marked := range h.Objects {
		if !marked {
			freed++
		}
	}
	return freed
}
```

## Walkthrough

`[true, false, true]` has one `false`, so `Sweep` returns 1. An empty slice
iterates zero times and returns the initial 0 — no special case needed.

## Pitfalls

- **Counting `true` instead.** Returns the live count, which is the complement
  of what was asked.
- **Reslicing or deleting during the walk.** Mutating the slice you are ranging
  over is a reliable source of skipped elements; a real sweep frees into a free
  list rather than shrinking the heap array.
- **Forgetting to clear the marks.** Not part of this puzzle, but a collector
  that does not reset the bits before the next mark phase will consider
  everything permanently live.

## Why Go's collector does it lazily

Go sweeps incrementally: spans are swept when they are next needed for
allocation, spreading the cost instead of stopping the world for a full heap
scan. The counting loop here is the conceptual version of that work.
