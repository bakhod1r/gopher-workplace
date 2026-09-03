# Read Through A Pointer That May Be Nil

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An optional field is modelled as `*int`. Every read site dereferences it, and the one that forgot the nil check is the one that runs in production.

## Task

Implement [derefsafe.go](derefsafe.go):

1. Return what `p` points at.
2. Return 0 when `p` is nil.
3. Zero allocations.

Replace the stub body in [derefsafe.go](derefsafe.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  n := 42; Value(&n)
Output: 42
```

**Example 2:**

```
Input:  Value(nil)
Output: 0
```

_Explanation:_ No panic.

**Example 3:**

```
Input:  n changed after taking &n
Output: Value sees the new value
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **nil is a valid pointer value** | Holding one is fine; dereferencing one is a panic. |
| 2 | **Pointers are live views** | They read the variable, not a snapshot of it. |
| 3 | **The zero value as a default** | 0 for a missing int is a decision, not an accident. |

## Hint

One comparison, then the dereference.

## Validate

```bash
make verify
```
