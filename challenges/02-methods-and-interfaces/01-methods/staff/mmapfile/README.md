# Mapped Region Reads

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A memory-mapped file is just a byte region; reading from it is a copy, not a
syscall. The interesting part is the contract at the edges — what a read that
runs off the end returns, and what an out-of-range offset returns — because
callers written against `io.ReaderAt` depend on it exactly.

## Task

Implement `ReadAt` on `*Mmap` in [mmapfile.go](mmapfile.go):

1. `off < 0` or `off >= len(m.Data)` returns `(0, io.EOF)`.
2. Copy as many bytes as fit into `p` starting at `off`.
3. If the copy reached the end of the region, return the count **and** `io.EOF`.
4. Otherwise return `(len(p), nil)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Data [10 20 30 40 50], p len 3, off 1
Output: (3, nil), p == [20 30 40]
```

**Example 2:**

```
Input:  Data [10 20 30], p len 4, off 1
Output: (2, io.EOF), p[:2] == [20 30]
```

**Example 3:**

```
Input:  Data [10 20 30], off 3 (or 99, or -1)
Output: (0, io.EOF)
```

_Explanation:_ a short read reports both the bytes it delivered and why it stopped.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`copy` returns the count** | `n := copy(p, m.Data[off:])` does the min-length arithmetic for you. |
| 2 | **Non-zero n with a non-nil error** | `io.ReaderAt` explicitly allows it; callers must use `n` before checking `err`. |
| 3 | **Bounds before slicing** | `m.Data[off:]` panics for a negative or oversized `off` — the guard must come first. |

## Hint

`copy` already stops at whichever slice is shorter, so the whole body is a
guard, one `copy`, and a decision about whether `n < len(p)` means EOF.

## Validate

```bash
make verify
```
