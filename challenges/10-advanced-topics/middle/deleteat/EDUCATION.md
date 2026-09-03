# Remove One Element, Release Its Slot

## Intuition

Shifting the tail down leaves the final slot holding a copy of the element before it. Shortening the slice hides that slot from you and not from the collector, so a long-lived array keeps one stale reference per deletion.

## Approach

1. Reject an out-of-range index.
2. `copy(s[i:], s[i+1:])` to close the gap.
3. Set the last slot to nil, then return `s[:len(s)-1]`.

## Solution

```go
// Item is one stored element.
type Item struct {
	ID  int
	Pad [512]byte
}

// DeleteAt removes the element at index i, preserving the order of the
// rest, and returns the shortened slice.
//
// The vacated slot at the end must be cleared so the removed item stops
// being reachable through the backing array.
//
// Examples:
//
// 	DeleteAt([]*Item{a, b, c}, 1) => a slice holding a and c
func DeleteAt(s []*Item, i int) []*Item {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	return s[:len(s)-1]
}
```

## Walkthrough

Deleting index 0 of [a b c] shifts b and c down, leaving c also at index 2. Clearing index 2 and returning `s[:2]` releases the duplicate.

## Pitfalls

- `append(s[:i], s[i+1:]...)` — the same shift, with the same uncleared tail.
- Clearing after reslicing, which no longer reaches the slot.
