# Iterator Pattern

## Intuition

An iterator is a cursor with two questions: *is there more?* and *give me the
next one*. Keeping the cursor in its own type means the collection stays
immutable and several traversals can run at once.

## Approach

1. `HasNext` is a pure read of the cursor.
2. `Next` snapshots the value, advances the cursor, and returns the snapshot.

## Solution

```go
func (it *ListIter) HasNext() bool {
	return it.current != nil
}

func (it *ListIter) Next() int {
	v := it.current.Val
	it.current = it.current.Next
	return v
}
```

## Walkthrough

`NewIter(head)` parks `current` on node 1. The test loop asks `HasNext()`
(true), then `Next()` returns 1 and moves `current` to node 2. After the third
`Next`, `current` is node 3's `Next`, which is nil, so `HasNext()` is false and
the loop ends with `[1 2 3]`.

## Pitfalls

- **Advancing before reading.** `it.current = it.current.Next; return
  it.current.Val` skips the first element and panics on the last.
- **Value receiver on `Next`.** The cursor never moves — an infinite loop that
  yields `1` forever.
- **Calling `Next` without `HasNext`.** The documented contract requires the
  check; without it, the nil `current` dereference panics.

## Two method sets, one name

`Node.Next` is a field; `ListIter.Next` is a method. They never collide because
they belong to different types — `it.current.Next` and `it.Next()` are resolved
against different declarations.
