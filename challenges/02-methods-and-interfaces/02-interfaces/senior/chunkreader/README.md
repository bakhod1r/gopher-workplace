# Chunked Reader

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A parser reads a 100M-line file. The current version reads it all into one string; the fix reads fixed-size chunks into a reused buffer.

## Task

Implement the stub(s) in [chunkreader.go](chunkreader.go):

1. Implement `Read` on `*ChunkSource`, filling the caller's buffer and returning how many bytes were written.
2. Implement `CountLines`, which counts `\n` bytes by streaming through a caller-supplied buffer.
3. Constraint: `CountLines` must allocate nothing per chunk — the buffer is supplied once and reused.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountLines(source("a\nb\n"), make([]byte, 4))
Output: 2
```

**Example 2:**

```
Input:  a chunk boundary splitting a line
Output: still counted once
```

**Example 3:**

```
Input:  an empty source
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Caller-supplied buffers** | The `io.Reader` contract: the caller owns the memory. |
| 2 | **Chunk boundaries** | State must survive across reads; a line can straddle a chunk. |
| 3 | **Zero allocation streaming** | Efficiency: reuse one buffer instead of allocating per chunk. |

## Hint

Copy into `p` with `copy(p, ...)` and return the copied count — never allocate inside `Read`.

## Validate

```bash
make verify
```
