# The Ring That Lends Out Its Buffer

**Level:** staff  
**Topic:** 03-generics

## Context

A telemetry ring buffer hands a snapshot to the reporting layer. The reporter appends a synthetic summary row to that snapshot, and the ring's next reading comes back as the summary row instead of a real sample.

## Task

Fix the single planted bug in [ringcapbug.go](ringcapbug.go):

1. Find and fix the single bug so the non-wrapped snapshot cannot be written through.
2. The element order and the wrapped case must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Push(1..3) on cap 8; Slice()
Output: [1 2 3] with cap 3
```

**Example 2:**

```
Input:  append to the snapshot, then Push(4)
Output: the snapshot keeps its appended value
```

**Example 3:**

```
Input:  Push(1..5) on cap 4; Slice()
Output: [2 3 4 5]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three-index slicing** | `s[a:b:b]` caps the result so `append` must allocate instead of overwriting. |
| 2 | **Length versus capacity** | A slice header carries spare capacity that `append` will happily write into. |
| 3 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate your internals. |

## Hint

Compare the capacity of the returned slice with its length.

## Validate

```bash
make verify
```
