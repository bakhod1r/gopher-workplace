# Bounds-Checked Index

**Level:** junior
**Topic:** 04-error-handling

## Context

A CSV importer reads columns by position. A short row must produce a clear error, not a runtime panic that kills the import.

## Task

Implement `At` in [elemat.go](elemat.go):

1. Return `s[i]` when `i` is within bounds.
2. Return `0` and `ErrOutOfRange` for any other index.
3. Reject negative indexes too.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  At([]int{1, 2, 3}, 0)
Output: 1, nil
```

**Example 2:**

```
Input:  At([]int{1, 2, 3}, 3)
Output: 0, ErrOutOfRange
```

**Example 3:**

```
Input:  At([]int{1}, -1)
Output: 0, ErrOutOfRange
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounds checking** | Valid indexes are `0 <= i < len(s)`. |
| 2 | **Panic vs error** | A guarded index turns a crash into a value. |
| 3 | **Boolean operators** | `||` combines the two invalid cases. |

## Hint

Two ways to be out of range — below zero and at or past the length.

## Validate

```bash
make verify
```
