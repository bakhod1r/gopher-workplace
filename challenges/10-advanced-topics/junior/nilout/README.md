# Let The Dropped Elements Be Collected

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A pool keeps a fixed-size slice of task pointers and reuses it between rounds. Memory grows round over round even though the tasks are finished.

## Task

Implement [nilout.go](nilout.go):

1. Set every element of `s` to nil.
2. Keep `len(s)` unchanged — the slice itself is reused.
3. A nil slice must not panic.

Replace the stub body in [nilout.go](nilout.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s := []*Node{{1},{2}}; DropAll(s)
Output: [<nil> <nil>]
```

**Example 2:**

```
Input:  len(s) after DropAll
Output: unchanged
```

_Explanation:_ Only the elements are released, not the slice.

**Example 3:**

```
Input:  DropAll(nil)
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reachability** | The collector frees what nothing points at; a stale pointer in a live slice is a pointer. |
| 2 | **clear on slices** | `clear(s)` writes the zero value into every element. |
| 3 | **Slices share storage** | Writing through the parameter is visible to the caller. |

## Hint

Reslicing to zero length hides the pointers from you, not from the collector.

## Validate

```bash
make verify
```
