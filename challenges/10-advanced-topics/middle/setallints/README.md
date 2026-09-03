# Write Every Int Field At Once

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A test helper resets every counter in a stats struct before each case. Adding a counter means remembering to reset it, and nobody does.

## Task

Implement [setallints.go](setallints.go):

1. Set every settable int field of the struct to `v`.
2. Return how many fields were written.
3. Skip unexported fields and every other kind.
4. Return `ErrTarget` unless `ptr` is a non-nil pointer to a struct.

Replace the stub body in [setallints.go](setallints.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  SetAllInts(&rec{}, 7)
Output: 2, nil
```

_Explanation:_ Only the two int fields.

**Example 2:**

```
Input:  a struct with no int fields
Output: 0, nil
```

**Example 3:**

```
Input:  SetAllInts(rec{}, 1)
Output: ErrTarget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **CanSet** | One check covers addressability and export status. |
| 2 | **Kind filtering** | `SetInt` panics on the wrong kind, so filter before writing. |
| 3 | **Counting the work** | Returning the count makes the helper testable without inspecting the struct. |

## Hint

`CanSet` is the only export check you need on the Value side.

## Validate

```bash
make verify
```
