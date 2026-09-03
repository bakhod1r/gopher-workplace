# Many Readers Over One Reinterpreted Buffer

**Level:** staff
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

An analytics kernel receives half a gigabyte of packed int64 records and wants every core summing at once, without first copying the buffer into a typed slice.

## Task

Implement [sharedview.go](sharedview.go):

1. Reinterpret `b` as `[]int64` and total it across `workers` goroutines over disjoint chunks.
2. Report false for an empty buffer, a length that is not a multiple of eight, or a misaligned start.
3. Any worker count must produce the same total; nothing may be counted twice.
4. The view must share `b`'s storage — no copy of the data.

Replace the stub body in [sharedview.go](sharedview.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  SumParallel(bytesOf([1,2,3,4]), 2)
Output: 10, true
```

**Example 2:**

```
Input:  workers = 100000 over 1001 values
Output: the same total, true
```

_Explanation:_ The worker count is clamped.

**Example 3:**

```
Input:  a 12-byte buffer
Output: 0, false
```

_Explanation:_ Not a multiple of 8.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reinterpreting a buffer** | `unsafe.Slice` over the data pointer with an element count. |
| 2 | **Concurrent reads are free** | Disjoint, read-only access needs no synchronisation. |
| 3 | **Chunk clamping** | The last chunk must stop at the view's length. |
| 4 | **Preconditions before parallelism** | Alignment and divisibility are decided once, up front. |

## Hint

Validate and build the view first; from there it is an ordinary parallel sum over a slice.

## Validate

```bash
make verify
```
