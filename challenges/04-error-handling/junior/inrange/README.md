# Bounded Value

**Level:** junior
**Topic:** 04-error-handling

## Context

A form takes a numeric setting with caller-supplied limits. The bounds themselves can be wrong, and that is its own failure.

## Task

Implement `InRange` in [inrange.go](inrange.go):

1. Return `ErrBadBounds` when `lo > hi`.
2. Return `ErrOutOfRange` when `n` falls outside `[lo, hi]`.
3. Return nil when `n` is inside the inclusive range.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  InRange(5, 1, 10)
Output: nil
```

**Example 2:**

```
Input:  InRange(0, 1, 10)
Output: ErrOutOfRange
```

**Example 3:**

```
Input:  InRange(5, 10, 1)
Output: ErrBadBounds
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Validating the validator** | Bad limits are checked before the value. |
| 2 | **Inclusive range** | `lo` and `hi` are both acceptable. |
| 3 | **Guard ordering** | The structural error takes priority. |

## Hint

Check the bounds before the value — otherwise a reversed range rejects everything with the wrong error.

## Validate

```bash
make verify
```
