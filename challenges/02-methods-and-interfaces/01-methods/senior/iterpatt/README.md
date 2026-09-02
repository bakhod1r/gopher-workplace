# Iterator Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A linked list has no index, so `for i := range` cannot walk it. An iterator
holds the traversal position as state and exposes two methods, so the caller
loops without knowing the list's shape.

## Task

Implement `HasNext` and `Next` on `*ListIter` in [iterpatt.go](iterpatt.go):

1. `HasNext()` reports whether `it.current` is non-nil.
2. `Next()` returns the current node's `Val` and advances `it.current` to `Next`.

**Constraint (senior):** a 1M-node list must be walked in a single pass holding only a cursor — no copy of the list, no per-step allocation.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  list 1→2→3; HasNext()
Output: true
```

**Example 2:**

```
Input:  Next(), Next(), Next()
Output: 1, 2, 3
```

**Example 3:**

```
Input:  HasNext() after the third Next()
Output: false
```

_Explanation:_ `current` has advanced past the tail to nil.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Iterator state** | The position lives in the iterator, not in the list — two iterators can walk the same list independently. |
| 2 | **Read-then-advance** | Save the value before moving the pointer, or you return the wrong element. |
| 3 | **Pointer receiver** | `Next` mutates `current`, so a value receiver would never advance. |

## Hint

`Next` is three lines: `v := it.current.Val`, `it.current = it.current.Next`,
`return v`. Note that the field `Next` and the method `Next` share a name on
different types — that is legal.

## Validate

```bash
make verify
```
