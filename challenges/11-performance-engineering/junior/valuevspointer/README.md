# Receivers Decide What Gets Copied

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A value receiver copies the whole struct on every call. For a two-field point that is free; for a struct carrying a 64-element buffer it is 512 bytes memcpy'd per call, and any mutation lands on the copy. The receiver choice is both a correctness decision and a performance one.

## Task

Implement all three methods in [valuevspointer.go](valuevspointer.go):

1. `Inc` adds one to the counter the caller holds.
2. `Value` returns the current count.
3. `IncCopy` takes the receiver by value, increments the copy, and returns it — leaving the original alone.

## Examples

**Example 1:**

```
Input:  c.Inc(); c.Inc(); c.Value()
Output: 2
```

**Example 2:**

```
Input:  c.IncCopy().Value(), then c.Value()
Output: 1, then 0
```

**Example 3:**

```
Input:  c = c.IncCopy().IncCopy().IncCopy(); c.Value()
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receivers copy** | The cost is the size of the struct, paid on every single call. |
| 2 | **Mutation needs a pointer** | A value receiver can only change a copy that is about to be discarded. |
| 3 | **Consistency within a type** | Mixing receiver kinds on one type confuses method sets and readers alike. |

## Topics used again

Methods, pointers, structs, method sets.

## Hint

`IncCopy` mutates `c` and returns it — that is legal precisely because `c` is already a copy.

## Validate

```bash
make verify
```
