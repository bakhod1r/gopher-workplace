# Batching Writer

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A downstream API charges per call. Writes are now buffered and flushed in batches of a fixed size.

## Task

Implement the stub(s) in [batchwriter.go](batchwriter.go):

1. Implement `Write` on `*BatchWriter`, flushing automatically when the batch is full.
2. Implement `Flush`, which sends any partial batch and is a no-op when empty.
3. Constraint: the buffer must be reused across batches (`buf[:0]`, not a fresh slice), and no record may be lost or duplicated.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  size 2; write a, b
Output: one flush of [a b]
```

**Example 2:**

```
Input:  a third write then Flush
Output: a second flush of [c]
```

**Example 3:**

```
Input:  Flush with an empty buffer
Output: no call downstream
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Batching** | Amortise a per-call cost across many records. |
| 2 | **Buffer reuse** | `buf[:0]` keeps the backing array; a fresh slice allocates per batch. |
| 3 | **Flush semantics** | Reused: partial batches must not be silently dropped. |

## Hint

The sink must receive a copy — the buffer is reused right after the flush returns.

## Validate

```bash
make verify
```
