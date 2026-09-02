# Dedup Stream

## Intuition

Exact dedup costs memory proportional to the number of distinct ids — unbounded by design. A sliding window trades some recall for a hard ceiling, which is the trade a 100M-event stream forces you to make.

## Approach

1. `ExactSet.Seen` returns true for a known id, otherwise records and returns false.
2. `WindowSet.Seen` does the same, but evicts `order[0]` from both structures once the window is full.
3. Handle `N <= 0` by recording nothing.
4. `Dedup` streams and counts the ids the set reported as unseen.

## Solution

```go
func (e *ExactSet) Seen(id string) bool {
	if e.seen[id] {
		return true
	}
	e.seen[id] = true
	return false
}

func (w *WindowSet) Seen(id string) bool {
	if w.seen[id] {
		return true
	}
	if w.N <= 0 {
		return false
	}
	if len(w.order) >= w.N {
		oldest := w.order[0]
		w.order = w.order[1:]
		delete(w.seen, oldest)
	}
	w.seen[id] = true
	w.order = append(w.order, id)
	return false
}

func Dedup(src Source, set SeenSet) int {
	unique := 0
	for {
		id, ok := src.Next()
		if !ok {
			return unique
		}
		if !set.Seen(id) {
			unique++
		}
	}
}
```

## Walkthrough

`WindowSet{N: 1}` over `[a b a]`: recording `b` evicts `a`, so the final `a` looks new — 3 uniques. With `N` 2 nothing is evicted and the count is 2.

## Pitfalls

- Evicting from the map but not from `order`, so the slice becomes the unbounded structure.
- Re-appending an id already in the window, which lets duplicates consume window slots.
- Buffering the stream to dedup it, which reintroduces the memory problem the window exists to solve.
