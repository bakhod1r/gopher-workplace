# Let The Dropped Elements Be Collected

## Intuition

A slice you keep alive keeps every pointer inside it alive too. Shortening the slice does not erase the elements past the new length — they are still in the array, still reachable, still pinning their nodes.

## Approach

1. Call `clear(s)` to write nil into every element.
2. Do not reslice or reassign — the caller shares this array.

## Solution

```go
// Node is one payload the slice points at.
type Node struct {
	ID int
}

// DropAll clears every element of s to nil, in place.
//
// The length of s must not change; only the pointers it holds are released
// so the nodes they referenced become unreachable.
//
// Examples:
//
// 	s := []*Node{{1}}; DropAll(s) => s[0] == nil
func DropAll(s []*Node) {
	clear(s)
}
```

## Walkthrough

Three nodes, three pointers. After `clear(s)` the array holds three nils, nothing else references the nodes, and the next collection reclaims them.

## Pitfalls

- `s = s[:0]` — the pointers are still in the backing array.
- `s = nil` — rebinds the local parameter only.
