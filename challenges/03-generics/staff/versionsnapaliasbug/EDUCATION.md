# The Snapshot That Keeps Changing

## Intuition

`s.live = s.snaps[id]` gives the two slice headers one backing array. Every later edit to the live list writes through into the stored snapshot, so the checkpoint drifts and can never be restored faithfully again.

## Approach

1. Reject an out-of-range id.
2. Allocate a fresh slice the length of the snapshot.
3. Copy the snapshot into it and install it as the live list.

## Solution

```go
func (s *Store[T]) Restore(id int) bool {
	if id < 0 || id >= len(s.snaps) {
		return false
	}
	snap := s.snaps[id]
	s.live = make([]T, len(snap))
	copy(s.live, snap)
	return true
}

func (s *Store[T]) Snapshot() int {
	c := make([]T, len(s.live))
	copy(c, s.live)
	s.snaps = append(s.snaps, c)
	return len(s.snaps) - 1
}

func (s *Store[T]) Append(v T) {
	s.live = append(s.live, v)
}

func (s *Store[T]) Set(i int, v T) bool {
	if i < 0 || i >= len(s.live) {
		return false
	}
	s.live[i] = v
	return true
}

func (s *Store[T]) Items() []T {
	out := make([]T, len(s.live))
	copy(out, s.live)
	return out
}
```

## Walkthrough

After the first `Restore`, `Set(0, 77)` writes into the snapshot's own array; the second `Restore` therefore yields `[77 2 3]`.

## Pitfalls

- Copying on `Snapshot` only, which protects the first checkpoint and no later one.
- Using `append(s.live[:0], snap...)` — that writes into whatever `live` already pointed at.
- Trusting that a full-slice re-slice (`snap[:]`) somehow detaches the storage; it does not.
