# Write A Field Through A Pointer

**Level:** junior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A test fixture builder wants to poke one field of a struct by name. The first attempt passed the struct by value and every write silently went nowhere — until reflect started panicking instead.

## Task

Implement [setfield.go](setfield.go):

1. Set the named int field of the struct `ptr` points at.
2. Return `ErrNotSettable` for a non-pointer, a nil pointer, a non-struct, a missing field, an unexported field, or a field that is not an int.

Replace the stub body in [setfield.go](setfield.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  SetInt(&counters{}, "Hits", 42)
Output: nil, Hits is 42
```

**Example 2:**

```
Input:  SetInt(counters{}, "Hits", 1)
Output: ErrNotSettable
```

_Explanation:_ A value, not a pointer.

**Example 3:**

```
Input:  SetInt(&counters{}, "hidden", 1)
Output: ErrNotSettable
```

_Explanation:_ Unexported fields are not settable.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Addressability** | Only a Value obtained through a pointer's `Elem` can be set. |
| 2 | **CanSet** | The check that covers both addressability and export status. |
| 3 | **Kind before Set** | `SetInt` panics on a non-int field, so verify the kind first. |

## Hint

`reflect.ValueOf(x)` is never settable. What makes a Value addressable?

## Validate

```bash
make verify
```
