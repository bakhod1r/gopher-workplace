# Keep The Capacity, Drop The Contents

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A line-oriented writer allocates a fresh `[]byte` for every record. The records are small, the record count is not, and the allocator shows up at the top of the profile.

## Task

Implement [resetbuf.go](resetbuf.go):

1. Return a slice of length 0 that still owns `buf`'s array and capacity.
2. A nil input must return an empty result without panicking.

Replace the stub body in [resetbuf.go](resetbuf.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Reset(make([]byte, 8, 64))
Output: len 0, cap 64
```

**Example 2:**

```
Input:  append(Reset(buf), 'z')
Output: writes into buf's existing array
```

_Explanation:_ No allocation for the next record.

**Example 3:**

```
Input:  Reset(nil)
Output: len 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Length vs capacity** | Length is what you can index; capacity is what you already own. |
| 2 | **Buffer reuse** | Resetting to `[:0]` is the standard reuse idiom. |
| 3 | **Zero-value slices** | `nil[:0]` is legal and yields an empty slice. |

## Hint

One expression. It is a reslice.

## Validate

```bash
make verify
```
