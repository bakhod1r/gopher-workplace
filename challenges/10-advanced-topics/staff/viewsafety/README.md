# A View The Caller Cannot Corrupt

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A framing layer hands each parser a view of the shared read buffer. One parser appends to its view, the append fits in the spare capacity, and the next frame in the buffer is silently rewritten.

## Task

Implement [viewsafety.go](viewsafety.go):

1. Return the `n` bytes of `b` starting at `off`, sharing the storage.
2. The result's capacity must equal its length, so an append cannot reach the following bytes.
3. Report false for a negative offset or length, or a window running past the end.
4. Zero allocations; a zero-length window is legal.

Replace the stub body in [viewsafety.go](viewsafety.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Window([]byte("abcdef"), 2, 3)
Output: "cde", true
```

**Example 2:**

```
Input:  append to the result
Output: does not touch b[5]
```

_Explanation:_ The capacity is exactly 3.

**Example 3:**

```
Input:  Window(b, 4, 3) on a 6-byte buffer
Output: nil, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three-index slicing** | `s[:n:n]` is what makes the capacity a boundary rather than a suggestion. |
| 2 | **Spare capacity is someone else's data** | A view into a shared buffer must not expose the rest of it. |
| 3 | **Explicit bounds checks** | `off+n > len(b)` is the check the runtime no longer performs for you. |
| 4 | **unsafe.Add for the start** | The offset is applied in pointer space. |

## Hint

Getting the right bytes is half the job. The other half is the third index.

## Validate

```bash
make verify
```
